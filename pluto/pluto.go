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
type Poster struct{}

func (Poster) RequestUrl() (string, bool) {
   return "https://service-concierge.clusters.pluto.tv/v1/wv/alt", true
}

func (Poster) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (Poster) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Poster) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (e EpisodeClip) Dash() (*url.URL, bool) {
   for _, s := range e.Sources {
      if s.Type == "DASH" {
         return &s.File.Url, true
      }
   }
   return nil, false
}

type EpisodeClip struct {
   Sources []struct {
      File Url
      Type string
   }
}

type Url struct {
   Url url.URL
}

func (u *Url) UnmarshalText(text []byte) error {
   return u.Url.UnmarshalBinary(text)
}

func (v Video) Clip() (*EpisodeClip, error) {
   req, err := http.NewRequest("", "http://api.pluto.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v2/episodes/")
      if v.Id != "" {
         b.WriteString(v.Id)
      } else {
         b.WriteString(v.Episode)
      }
      b.WriteString("/clips.json")
      return b.String()
   }()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   var clips []EpisodeClip
   err = json.NewDecoder(resp.Body).Decode(&clips)
   if err != nil {
      return nil, err
   }
   return &clips[0], nil
}

var Base = []string{
   // these return `403 OK` with compressed content
   "http://siloh-fs.plutotv.net",
   "http://siloh-ns1.plutotv.net",
   "https://siloh-fs.plutotv.net",
   "https://siloh-ns1.plutotv.net",
   // returns `200 OK` with plain content
   "http://silo-hybrik.pluto.tv.s3.amazonaws.com",
}
