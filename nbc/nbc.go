package nbc

import (
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

func Core() Video {
   var v Video
   v.DrmProxySecret = "Whn8QFuLFM7Heiz6fYCYga7cYPM8ARe6"
   v.DrmProxyUrl = func() string {
      var b strings.Builder
      b.WriteString("https://drmproxy.digitalsvc.apps.nbcuni.com")
      b.WriteString("/drm-proxy/license")
      return b.String()
   }()
   return v
}

func (v Video) RequestUrl() (string, bool) {
   t, h := func() (int64, []byte) {
      h := hmac.New(sha256.New, []byte(v.DrmProxySecret))
      t := time.Now().UnixMilli()
      fmt.Fprint(h, t, "widevine")
      return t, h.Sum(nil)
   }()
   b := []byte(v.DrmProxyUrl)
   b = append(b, "/widevine"...)
   b = fmt.Append(b, "?time=", t)
   b = fmt.Appendf(b, "&hash=%x", h)
   b = append(b, "&device=web"...)
   return string(b), true
}

func (Video) RequestHeader() (http.Header, error) {
   head := make(http.Header)
   head.Set("content-type", "application/octet-stream")
   return head, nil
}

type Video struct {
   DrmProxyUrl string
   DrmProxySecret string
}

func (m Metadata) OnDemand() (*OnDemand, error) {
   req, err := http.NewRequest("GET", "https://lemonade.nbc.com", nil)
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
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   video := new(OnDemand)
   err = json.NewDecoder(res.Body).Decode(video)
   if err != nil {
      return nil, err
   }
   return video, nil
}

func (Video) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Video) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

type OnDemand struct {
   PlaybackUrl string
}

const query = `
query bonanzaPage(
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
