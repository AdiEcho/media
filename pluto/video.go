package pluto

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "slices"
   "strconv"
   "strings"
)

var Forward string

func (w WebAddress) Video() (*Video, error) {
   req, err := http.NewRequest("GET", "https://boot.pluto.tv/v4/start", nil)
   if err != nil {
      return nil, err
   }
   if Forward != "" {
      req.Header.Set("x-forwarded-for", Forward)
   }
   req.URL.RawQuery = url.Values{
      "appName":           {"web"},
      "appVersion":        {"9"},
      "clientID":          {"9"},
      "clientModelNumber": {"9"},
      "drmCapabilities":   {"widevine:L3"},
      "seriesIDs":         {w.series},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var start struct {
      VOD []Video
   }
   err = json.NewDecoder(res.Body).Decode(&start)
   if err != nil {
      return nil, err
   }
   demand := start.VOD[0]
   if demand.Slug.Slug != w.series {
      if demand.ID != w.series {
         return nil, errors.New(demand.Slug.Slug)
      }
   }
   for _, s := range demand.Seasons {
      s.parent = &demand
      for _, e := range s.Episodes {
         err := e.Slug.atoi()
         if err != nil {
            return nil, err
         }
         e.parent = s
         if e.Episode == w.episode {
            return e, nil
         }
         if e.Slug.Slug == w.episode {
            return e, nil
         }
      }
   }
   err = demand.Slug.atoi()
   if err != nil {
      return nil, err
   }
   return &demand, nil
}

// ex-machina-2015-1-1-ptv1
// head-first-1998-1-2
// king-of-queens
// pilot-1998-1-1-ptv8
func (s *Slug) atoi() error {
   split := strings.Split(s.Slug, "-")
   slices.Reverse(split)
   if strings.HasPrefix(split[0], "ptv") {
      split = split[1:]
   }
   var err error
   s.episode, err = strconv.Atoi(split[0])
   if err != nil {
      return err
   }
   s.season, err = strconv.Atoi(split[1])
   if err != nil {
      return err
   }
   s.year, err = strconv.Atoi(split[2])
   if err != nil {
      return err
   }
   return nil
}

func (s *Slug) UnmarshalText(text []byte) error {
   s.Slug = string(text)
   return nil
}

type Season struct {
   Episodes []*Video
   parent   *Video
}

func (n Namer) Show() string {
   if v := n.V.parent; v != nil {
      return v.parent.Name
   }
   return ""
}

type Namer struct {
   V *Video
}

func (n Namer) Season() int {
   return n.V.Slug.season
}

func (n Namer) Episode() int {
   return n.V.Slug.episode
}

func (n Namer) Title() string {
   return n.V.Name
}

func (n Namer) Year() int {
   return n.V.Slug.year
}
func (w WebAddress) String() string {
   var b strings.Builder
   if w.series != "" {
      b.WriteString("https://pluto.tv/on-demand/")
      if w.episode != "" {
         b.WriteString("series")
      } else {
         b.WriteString("movies")
      }
      b.WriteByte('/')
      b.WriteString(w.series)
   }
   if w.episode != "" {
      b.WriteString("/episode/")
      b.WriteString(w.episode)
   }
   return b.String()
}

type WebAddress struct {
   series  string
   episode string
}

type Video struct {
   Name    string
   Seasons []*Season
   parent  *Season
   Slug    Slug
   Episode string `json:"_id"`
   ID      string
}

type Slug struct {
   episode int
   season  int
   Slug    string
   year    int
}

func (w *WebAddress) Set(s string) error {
   for {
      var (
         key string
         ok  bool
      )
      key, s, ok = strings.Cut(s, "/")
      if !ok {
         return nil
      }
      switch key {
      case "episode":
         w.episode = s
      case "movies":
         w.series = s
      case "series":
         w.series, s, ok = strings.Cut(s, "/")
         if !ok {
            return errors.New("episode")
         }
      }
   }
}
