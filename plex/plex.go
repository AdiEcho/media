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

func (a anonymous) abs(path string, query url.Values) string {
   query.Set("x-plex-token", a.AuthToken)
   var u url.URL
   u.Host = "vod.provider.plex.tv"
   u.Path = path
   u.RawQuery = query.Encode()
   u.Scheme = "https"
   return u.String()
}

type anonymous struct {
   AuthToken string
}

func (a *anonymous) New() error {
   req, err := http.NewRequest(
      "POST", "https://plex.tv/api/v2/users/anonymous", nil,
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Accept": {"application/json"},
      "X-Plex-Product": {"Plex Mediaverse"},
      "X-Plex-Client-Identifier": {"!"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
