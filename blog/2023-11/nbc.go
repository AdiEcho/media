package nbc

import (
   "net/http"
   "net/url"
)

type Hello struct{}

func (Hello) Request_URL() string {
   var u url.URL
   u.Host = "drmproxy.digitalsvc.apps.nbcuni.com"
   u.Path = "/drm-proxy/license/widevine"
   u.RawQuery = "device=web"
   u.Scheme = "https"
   return u.String()
}

func (Hello) Request_Header() http.Header {
   return nil
}

func (Hello) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (Hello) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}
