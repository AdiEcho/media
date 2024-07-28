package main

import (
   "crypto/hmac"
   "crypto/sha256"
   "fmt"
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

type Video struct {
   DrmProxyUrl string
   DrmProxySecret string
}

func (v Video) RequestUrl() (string, bool) {
   t, h := func() (int64, []byte) {
      h := hmac.New(sha256.New, []byte(v.DrmProxySecret))
      t := time.Now().UnixMilli()
      
      t = 1722186920760
      
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

func main() {
   address, _ := Core().RequestUrl()
   println(address)
}
