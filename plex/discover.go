package plex

import (
   "encoding/json"
   "net/http"
   "net/url"
)

func (a anonymous) matches(path string) (*metadata, error) {
   req, err := http.NewRequest(
      "GET", "https://discover.provider.plex.tv/library/metadata/matches", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("accept", "application/json")
   req.URL.RawQuery = url.Values{
      "url": {path},
      "x-plex-token": {a.AuthToken},
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

func (metadata) Show() string {
   return ""
}

func (metadata) Season() int {
   return 0
}

func (metadata) Episode() int {
   return 0
}

func (metadata) Title() string {
   return ""
}

func (metadata) Year() int {
   return 0
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
type metadata struct {
   Media []struct {
      Part []part
      Protocol string
   }
   RatingKey string
}
