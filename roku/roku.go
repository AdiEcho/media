package roku

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

type Playback struct {
   DRM struct {
      Widevine struct {
         LicenseServer string
      }
   }
}

func (c CrossSite) csrf() (*http.Cookie, bool) {
   for _, cookie := range c.cookies {
      if cookie.Name == "_csrf" {
         return cookie, true
      }
   }
   return nil, false
}

func (Playback) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (p Playback) RequestUrl() (string, bool) {
   return p.DRM.Widevine.LicenseServer, true
}

func (Playback) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}
