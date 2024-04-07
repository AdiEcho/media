package plex

import (
   "encoding/json"
   "net/http"
   "net/url"
)

func (discover) Show() string {
   return ""
}

func (discover) Season() int {
   return 0
}

func (discover) Episode() int {
   return 0
}

func (discover) Title() string {
   return ""
}

func (discover) Year() int {
   return 0
}

type discover struct {
   RatingKey string
}

func (a anonymous) discover(path string) (*discover, error) {
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
         Metadata []discover
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   return &s.MediaContainer.Metadata[0], nil
}
