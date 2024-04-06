package plex

import (
   "net/http"
   "net/url"
)

func (part) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (part) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (part) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

type part struct {
   Key string
   License string
}

func (p part) RequestUrl() (string, bool) {
   return p.License, true
}

func (m metadata) dash(a anonymous) (*part, bool) {
   for _, media := range m.Media {
      if media.Protocol == "dash" {
         p := media.Part[0]
         p.Key = a.abs(p.Key, url.Values{})
         p.License = a.abs(p.License, url.Values{
            "x-plex-drm": {"widevine"},
         })
         return &p, true
      }
   }
   return nil, false
}
