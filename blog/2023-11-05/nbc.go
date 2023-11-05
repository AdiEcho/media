package nbc

import (
   "crypto/hmac"
   "crypto/sha256"
   "fmt"
   "net/http"
   "net/url"
   "time"
)

type Hello struct {
   Core_Video struct {
      DRM_Proxy_Secret string `json:"drmProxySecret"`
      DRM_Proxy_URL string `json:"drmProxyUrl"`
   } `json:"coreVideo"`
}

func (h Hello) Request_URL() string {
   var u url.URL
   u.Host = "drmproxy.digitalsvc.apps.nbcuni.com"
   u.Path = "/drm-proxy/license/widevine"
   u.Scheme = "https"
   u.RawQuery = func() string {
      t := time.Now().UnixMilli()
      v := make(url.Values)
      v.Set("device", "web")
      v.Set("time", fmt.Sprint(t))
      w := hmac.New(sha256.New, []byte(h.Core_Video.DRM_Proxy_Secret))
      fmt.Fprint(w, t, "widevine")
      v.Set("hash", fmt.Sprintf("%x", w.Sum(nil)))
      return v.Encode()
   }()
   return u.String()
}

func (Hello) Request_Header() http.Header {
   h := make(http.Header)
   h["Content-Type"] = []string{"application/octet-stream"}
   return h
}

func (Hello) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (Hello) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}
