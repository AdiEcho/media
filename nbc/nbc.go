package nbc

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/json"
   "errors"
   "fmt"
   "net/http"
   "net/url"
   "strconv"
   "strings"
   "time"
)

const query = `
query(
   $app: NBCUBrands!
   $name: String!
   $oneApp: Boolean
   $platform: SupportedPlatforms!
   $type: EntityPageType!
   $userId: String!
) {
   bonanzaPage(
      app: $app
      name: $name
      oneApp: $oneApp
      platform: $platform
      type: $type
      userId: $userId
   ) {
      metadata {
         ... on VideoPageData {
            airDate
            episodeNumber
            movieShortTitle
            mpxAccountId
            mpxGuid
            programmingType
            seasonNumber
            secondaryTitle
            seriesShortTitle
         }
      }
   }
}
`

// this is better than strings.Replace and strings.ReplaceAll
func graphql_compact(s string) string {
   field := strings.Fields(s)
   return strings.Join(field, " ")
}

type OnDemand struct {
   PlaybackUrl string
}

type page_request struct {
   Query string `json:"query"`
   Variables struct {
      App string `json:"app"` // String cannot represent a non string value
      Name string `json:"name"`
      OneApp bool `json:"oneApp"`
      Platform string `json:"platform"`
      Type string `json:"type"` // can be empty
      UserId string `json:"userId"`
   } `json:"variables"`
}

type Metadata struct {
   AirDate time.Time
   EpisodeNumber int `json:",string"`
   MovieShortTitle string
   MpxAccountId int64 `json:",string"`
   MpxGuid int64 `json:",string"`
   ProgrammingType string
   SeasonNumber int `json:",string"`
   SecondaryTitle string
   SeriesShortTitle string
}

func (m *Metadata) Show() string {
   return m.SeriesShortTitle
}

func (m *Metadata) Season() int {
   return m.SeasonNumber
}

func (m *Metadata) Episode() int {
   return m.EpisodeNumber
}

func (m *Metadata) Year() int {
   return m.AirDate.Year()
}

func (m *Metadata) Title() string {
   if m.MovieShortTitle != "" {
      return m.MovieShortTitle
   }
   return m.SecondaryTitle
}

func (m *Metadata) OnDemand() (*OnDemand, error) {
   req, err := http.NewRequest("", "https://lemonade.nbc.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/v1/vod/")
      b = strconv.AppendInt(b, m.MpxAccountId, 10)
      b = append(b, '/')
      b = strconv.AppendInt(b, m.MpxGuid, 10)
      return string(b)
   }()
   req.URL.RawQuery = url.Values{
      "platform": {"web"},
      "programmingType": {m.ProgrammingType},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   video := &OnDemand{}
   err = json.NewDecoder(resp.Body).Decode(video)
   if err != nil {
      return nil, err
   }
   return video, nil
}

func (m *Metadata) New(guid int) error {
   data, err := func() ([]byte, error) {
      var p page_request
      p.Query = graphql_compact(query)
      p.Variables.App = "nbc"
      p.Variables.Name = strconv.Itoa(guid)
      p.Variables.OneApp = true
      p.Variables.Platform = "android"
      p.Variables.Type = "VIDEO"
      return json.Marshal(p)
   }()
   if err != nil {
      return err
   }
   resp, err := http.Post(
      "https://friendship.nbc.co/v2/graphql", "application/json",
      bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   var value struct {
      Data struct {
         BonanzaPage struct {
            Metadata Metadata
         }
      }
      Errors []struct {
         Message string
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return err
   }
   if v := value.Errors; len(v) >= 1 {
      return errors.New(v[0].Message)
   }
   *m = value.Data.BonanzaPage.Metadata
   return nil
}

///

type Video struct {
   DrmProxyUrl string
   DrmProxySecret string
}

func (v *Video) RequestUrl() (string, bool) {
   now := time.Now().UnixMilli()
   mac := hmac.New(sha256.New, []byte(v.DrmProxySecret))
   fmt.Fprint(mac, now, "widevine")
   b := []byte(v.DrmProxyUrl)
   b = append(b, "/widevine"...)
   b = append(b, "?device=web"...)
   b = fmt.Append(b, "&time=", now)
   b = fmt.Appendf(b, "&hash=%x", mac.Sum(nil))
   return string(b), true
}

func (Video) RequestHeader() (http.Header, error) {
   head := http.Header{}
   head.Set("content-type", "application/octet-stream")
   return head, nil
}

func (Video) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Video) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (v *Video) New() {
   v.DrmProxySecret = "Whn8QFuLFM7Heiz6fYCYga7cYPM8ARe6"
   v.DrmProxyUrl = func() string {
      var b strings.Builder
      b.WriteString("https://drmproxy.digitalsvc.apps.nbcuni.com")
      b.WriteString("/drm-proxy/license")
      return b.String()
   }()
}
