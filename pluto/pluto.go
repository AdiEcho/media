package pluto

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (a Address) Video(forward string) (*OnDemand, error) {
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
   var data struct {
      Vod []OnDemand
   }
   err = json.NewDecoder(resp.Body).Decode(&data)
   if err != nil {
      return nil, err
   }
   demand := data.Vod[0]
   if demand.Slug.Slug != a.series {
      if demand.Id != a.series {
         return nil, errors.New(demand.Slug.Slug)
      }
   }
   for _, s := range demand.Seasons {
      s.show = &demand
      for _, episode := range s.Episodes {
         err := episode.Slug.atoi()
         if err != nil {
            return nil, err
         }
         episode.season = s
         if episode.Episode == a.episode {
            return episode, nil
         }
         if episode.Slug.Slug == a.episode {
            return episode, nil
         }
      }
   }
   err = demand.Slug.atoi()
   if err != nil {
      return nil, err
   }
   return &demand, nil
}

var Base = []struct{
   Scheme string
   Host string
   Status string
}{
   {"https", "siloh-fs.plutotv.net", "403 OK"},
   {"https", "siloh-ns1.plutotv.net", "403 OK"},
   {"http", "siloh-fs.plutotv.net", "403 OK"},
   {"http", "siloh-ns1.plutotv.net", "403 OK"},
   {"http", "silo-hybrik.pluto.tv.s3.amazonaws.com", "200 OK"},
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

type Address struct {
   series  string
   episode string
}

func (e *EpisodeClip) Dash() (*Url, bool) {
   for _, source := range e.Sources {
      if source.Type == "DASH" {
         return &source.File, true
      }
   }
   return nil, false
}

func (a *Address) String() string {
   var b strings.Builder
   if a.series != "" {
      if a.episode != "" {
         b.WriteString("series/")
         b.WriteString(a.series)
         b.WriteString("/episode/")
         b.WriteString(a.episode)
      } else {
         b.WriteString("movies/")
         b.WriteString(a.series)
      }
   }
   return b.String()
}

type EpisodeClip struct {
   Sources []struct {
      File Url
      Type string
   }
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

func (o OnDemand) Clip() (*EpisodeClip, error) {
   req, err := http.NewRequest("", "http://api.pluto.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b strings.Builder
      b.WriteString("/v2/episodes/")
      if o.Id != "" {
         b.WriteString(o.Id)
      } else {
         b.WriteString(o.Episode)
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

type Url struct {
   Url *url.URL
}

func (u *Url) UnmarshalText(text []byte) error {
   u.Url = &url.URL{}
   return u.Url.UnmarshalBinary(text)
}

type Season struct {
   Episodes []*OnDemand
   show   *OnDemand
}

type Namer struct {
   Video *OnDemand
}

type OnDemand struct {
   Episode string `json:"_id"`
   Id      string
   Name    string
   Seasons []*Season
   Slug    string
   season  *Season
}

func (Namer) Year() int {
   return 0
}

func (Namer) Season() int {
   return 0
}

func (Namer) Episode() int {
   return 0
}

func (n Namer) Show() string {
   if v := n.Video.season; v != nil {
      return v.show.Name
   }
   return ""
}

func (n Namer) Title() string {
   return n.Video.Name
}
