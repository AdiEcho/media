package plex

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

type Path struct {
   s string
}

// https://watch.plex.tv/movie/the-hurt-locker
// https://watch.plex.tv/watch/movie/the-hurt-locker
func (p *Path) Set(s string) error {
   s = strings.TrimPrefix(s, "https://watch.plex.tv")
   p.s = strings.TrimPrefix(s, "/watch")
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
   var match struct {
      MediaContainer struct {
         Metadata []DiscoverMatch
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&match); err != nil {
      return nil, err
   }
   return &match.MediaContainer.Metadata[0], nil
}

type DiscoverMatch struct {
   GrandparentTitle string
   Index int
   ParentIndex int
   RatingKey string
   Title string
   Year int
}

type Namer struct {
   D *DiscoverMatch
}

func (n Namer) Episode() int {
   return n.D.Index
}

func (n Namer) Season() int {
   return n.D.ParentIndex
}

func (n Namer) Show() string {
   return n.D.GrandparentTitle
}

func (n Namer) Title() string {
   return n.D.Title
}

func (n Namer) Year() int {
   return n.D.Year
}
