package plex

import (
   "encoding/json"
   "net/http"
   "net/url"
)

func (a anonymous) discover(path string) (*discover_match, error) {
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
         Metadata []discover_match
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   return &s.MediaContainer.Metadata[0], nil
}

func (d discover_match) Show() string {
   return d.V.GrandparentTitle
}

func (d discover_match) Season() int {
   return d.V.ParentIndex
}

func (d discover_match) Episode() int {
   return d.V.Index
}

func (d discover_match) Title() string {
   return d.V.Title
}

type discover_match struct {
   V struct {
      GrandparentTitle string
      Index int
      ParentIndex int
      RatingKey string
      Title string
      Year int
   }
}

func (d discover_match) Year() int {
   return d.V.Year
}
