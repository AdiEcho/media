package plex

import (
   "encoding/json"
   "net/http"
   "net/url"
)

func (m metadata) dash(auth_token string) (*part, bool) {
   for _, each := range m.Media {
      if each.Protocol == "dash" {
         p := each.Part[0]
         p.auth_token = auth_token
         return &p, true
      }
   }
   return nil, false
}

func (a anonymous) metadata(address string) (*metadata, error) {
   match, err := url.Parse(address)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "GET", "https://discover.provider.plex.tv/library/metadata/matches", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("accept", "application/json")
   req.URL.RawQuery = url.Values{
      "X-Plex-Token": {a.AuthToken},
      "url": {match.Path},
   }.Encode()
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

type part struct {
   Key string
   License string
   auth_token string
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

func (p part) RequestUrl() (string, bool) {
   var u url.URL
   u.Host = "vod.provider.plex.tv"
   u.Path = p.License
   u.Scheme = "https"
   u.RawQuery = url.Values{
      "X-Plex-DRM": {"widevine"},
      "X-Plex-Token": {p.auth_token},
   }.Encode()
   return u.String(), true
}
