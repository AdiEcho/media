package plex

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type Path struct {
   s string
}

func (p *Path) Set(s string) error {
   u, err := url.Parse(s)
   if err != nil {
      return err
   }
   p.s = u.Path
   return nil
}

func (p Path) String() string {
   return p.s
}

func (a Anonymous) Discover(p Path) (*DiscoverMatch, error) {
   req, err := http.NewRequest(
      "GET", "https://discover.provider.plex.tv/library/metadata/matches", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("accept", "application/json")
   req.URL.RawQuery = url.Values{
      "url": {p.s},
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
   var out DiscoverMatch
   out.m = in.MediaContainer.Metadata[0]
   return &out, nil
}

type DiscoverMatch struct {
   m metadata
}

func (d DiscoverMatch) Episode() int {
   return d.m.Index
}

func (d DiscoverMatch) Season() int {
   return d.m.ParentIndex
}

func (d DiscoverMatch) Show() string {
   return d.m.GrandparentTitle
}

func (d DiscoverMatch) Title() string {
   return d.m.Title
}

func (d DiscoverMatch) Year() int {
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
