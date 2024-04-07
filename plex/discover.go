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
   var in struct {
      MediaContainer struct {
         Metadata []metadata
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&in); err != nil {
      return nil, err
   }
   var out discover_match
   out.m = in.MediaContainer.Metadata[0]
   return &out, nil
}

type discover_match struct {
   m metadata
}

func (d discover_match) Episode() int {
   return d.m.Index
}

func (d discover_match) Season() int {
   return d.m.ParentIndex
}

func (d discover_match) Show() string {
   return d.m.GrandparentTitle
}

func (d discover_match) Title() string {
   return d.m.Title
}

func (d discover_match) Year() int {
   return d.m.Year
}

type metadata struct {
   GrandparentTitle string
   Index int
   ParentIndex int
   RatingKey string
   Title string
   Year int
}
