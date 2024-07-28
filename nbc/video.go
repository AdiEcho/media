package nbc

import (
   "crypto/hmac"
   "crypto/sha256"
   "fmt"
   "net/http"
   "strings"
   "time"
)

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

func (Video) RequestHeader() (http.Header, error) {
   head := make(http.Header)
   head.Set("content-type", "application/octet-stream")
   return head, nil
}

type Video struct {
   DrmProxyUrl string
   DrmProxySecret string
}

func (Video) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Video) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}
