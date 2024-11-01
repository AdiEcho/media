package nbc

import (
   "crypto/hmac"
   "crypto/sha256"
   "fmt"
   "net/http"
   "strings"
   "time"
)

// NO ANONYMOUS QUERY
const bonanza_page = `
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

type CoreVideo struct {
   DrmProxyUrl string
   DrmProxySecret string
}

func (c *CoreVideo) RequestUrl() (string, bool) {
   now := time.Now().UnixMilli()
   mac := hmac.New(sha256.New, []byte(c.DrmProxySecret))
   fmt.Fprint(mac, now, "widevine")
   b := []byte(c.DrmProxyUrl)
   b = append(b, "/widevine"...)
   b = append(b, "?device=web"...)
   b = fmt.Append(b, "&time=", now)
   b = fmt.Appendf(b, "&hash=%x", mac.Sum(nil))
   return string(b), true
}

func (c *CoreVideo) New() {
   c.DrmProxySecret = "Whn8QFuLFM7Heiz6fYCYga7cYPM8ARe6"
   c.DrmProxyUrl = func() string {
      var b strings.Builder
      b.WriteString("https://drmproxy.digitalsvc.apps.nbcuni.com")
      b.WriteString("/drm-proxy/license")
      return b.String()
   }()
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

func (*CoreVideo) RequestHeader() (http.Header, error) {
   head := http.Header{}
   head.Set("content-type", "application/octet-stream")
   return head, nil
}

func (*CoreVideo) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (*CoreVideo) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}
