package nbc

import (
   "crypto/hmac"
   "crypto/sha256"
   "fmt"
   "net/http"
   "time"
)

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

func (v Video) Request_URL() string {
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
   return string(b)
}

func (Video) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}
