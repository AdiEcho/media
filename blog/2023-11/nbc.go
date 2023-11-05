package nbc

import "net/http"

type Hello struct{}

func (Hello) Request_URL() string {
   return "https://drmproxy.digitalsvc.apps.nbcuni.com/drm-proxy/license/widevine"
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
