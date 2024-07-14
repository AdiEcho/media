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
   // some items just dont have a date:
   // bound-paramount-1-1
   // not just missing from the slug, missing EVERYWHERE, both on web and
   // Android
   s.year, _ = strconv.Atoi(split[2])
   return nil
}

func (a Address) Video(forward string) (*Video, error) {
   req, err := http.NewRequest("", "https://boot.pluto.tv/v4/start", nil)
   if err != nil {
      return nil, err
   }
   if forward != "" {
      req.Header.Set("x-forwarded-for", forward)
   }
   req.URL.RawQuery = url.Values{
      "appName":           {"web"},
      "appVersion":        {"9"},
      "clientID":          {"9"},
      "clientModelNumber": {"9"},
      "drmCapabilities":   {"widevine:L3"},
      "seriesIDs":         {a.series},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var start struct {
      Vod []Video
   }
   err = json.NewDecoder(resp.Body).Decode(&start)
   if err != nil {
      return nil, err
   }
   demand := start.Vod[0]
   if demand.Slug.Slug != a.series {
      if demand.Id != a.series {
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
         if e.Episode == a.episode {
            return e, nil
         }
         if e.Slug.Slug == a.episode {
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
func (a Address) String() string {
   var b strings.Builder
   if a.series != "" {
      b.WriteString("https://pluto.tv/on-demand/")
      if a.episode != "" {
         b.WriteString("series")
      } else {
         b.WriteString("movies")
      }
      b.WriteByte('/')
      b.WriteString(a.series)
   }
   if a.episode != "" {
      b.WriteString("/episode/")
      b.WriteString(a.episode)
   }
   return b.String()
}

type Address struct {
   series  string
   episode string
}

type Video struct {
   Name    string
   Seasons []*Season
   parent  *Season
   Slug    Slug
   Episode string `json:"_id"`
   Id      string
}

type Slug struct {
   episode int
   season  int
   Slug    string
   year    int
}

func (a *Address) Set(text string) error {
   for {
      var (
         key string
         ok  bool
      )
      key, text, ok = strings.Cut(text, "/")
      if !ok {
         return nil
      }
      switch key {
      case "episode":
         a.episode = text
      case "movies":
         a.series = text
      case "series":
         a.series, text, ok = strings.Cut(text, "/")
         if !ok {
            return errors.New("episode")
         }
      }
   }
}
