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

func (m Metadata) On_Demand() (*On_Demand, error) {
   req, err := http.NewRequest("GET", "https://lemonade.nbc.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/v1/vod/")
      b = strconv.AppendInt(b, m.MPX_Account_ID, 10)
      b = append(b, '/')
      b = strconv.AppendInt(b, m.MPX_GUID, 10)
      return string(b)
   }()
   req.URL.RawQuery = url.Values{
      "platform": {"web"},
      "programmingType": {m.Programming_Type},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   video := new(On_Demand)
   if err := json.NewDecoder(res.Body).Decode(video); err != nil {
      return nil, err
   }
   return video, nil
}

func (v Video) Request_URL() (string, error) {
   t, h := func() (int64, []byte) {
      h := hmac.New(sha256.New, []byte(v.DRM_Proxy_Secret))
      t := time.Now().UnixMilli()
      fmt.Fprint(h, t, "widevine")
      return t, h.Sum(nil)
   }()
   b := []byte(v.DRM_Proxy_URL)
   b = append(b, "/widevine"...)
   b = fmt.Append(b, "?time=", t)
   b = fmt.Appendf(b, "&hash=%x", h)
   b = append(b, "&device=web"...)
   return string(b), nil
}

type Video struct {
   DRM_Proxy_Secret string
   DRM_Proxy_URL string
}

var Core = Video{
   "Whn8QFuLFM7Heiz6fYCYga7cYPM8ARe6",
   "https://drmproxy.digitalsvc.apps.nbcuni.com/drm-proxy/license",
}

func (Video) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (Video) Request_Header() http.Header {
   return http.Header{
      "Content-Type": {"application/octet-stream"},
   }
}

func (Video) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}

type On_Demand struct {
   Playback_URL string `json:"playbackUrl"`
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

// this is better than strings.Replace and strings.ReplaceAll
func graphQL_compact(s string) string {
   f := strings.Fields(s)
   return strings.Join(f, " ")
}

type page_request struct {
   Query string `json:"query"`
   Variables struct {
      App string `json:"app"` // String cannot represent a non string value
      Name string `json:"name"`
      One_App bool `json:"oneApp"`
      Platform string `json:"platform"`
      Type string `json:"type"` // can be empty
      User_ID string `json:"userId"`
   } `json:"variables"`
}
