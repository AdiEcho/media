package plex

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (a anonymous) metadata(w web_address) (*metadata, error) {
   req, err := http.NewRequest("GET", "https://vod.provider.plex.tv", nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Accept": {"application/json"},
      "X-Plex-Token": {a.AuthToken},
   }
   req.URL.Path = "/library/metadata/" + w.String()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var s struct {
      MediaContainer struct {
         Metadata []metadata
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   return &s.MediaContainer.Metadata[0], nil
}

type metadata struct {
   Media []struct {
      Part []part
      Protocol string
   }
}

func (part) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (part) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (part) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

type web_address struct {
   key string
   value string
}

func (w *web_address) Set(s string) error {
   if i := strings.LastIndexByte(s, '/'); i >= 0 {
      s, w.value = s[:i], s[i+1:]
      if i := strings.LastIndexByte(s, '/'); i >= 0 {
         w.key = s[i+1:]
         return nil
      }
   }
   return errors.New("web_address.Set")
}

func (w web_address) String() string {
   var b strings.Builder
   b.WriteString(w.key)
   b.WriteByte(':')
   b.WriteString(w.value)
   return b.String()
}

type part struct {
   Key string
   License string
}

func (p part) RequestUrl() (string, bool) {
   return p.License, true
}

func (a anonymous) abs(path string, query url.Values) string {
   query.Set("x-plex-token", a.AuthToken)
   var u url.URL
   u.Host = "vod.provider.plex.tv"
   u.Path = path
   u.RawQuery = query.Encode()
   u.Scheme = "https"
   return u.String()
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
