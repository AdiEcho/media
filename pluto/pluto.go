package pluto

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

type EpisodeClip struct {
   Sources []struct {
      File Url
      Type string
   }
}

func (e *EpisodeClip) Dash() (*url.URL, bool) {
   for _, source := range e.Sources {
      if source.Type == "DASH" {
         return source.File.Url, true
      }
   }
   return nil, false
}

var Base = []FileBase{
   {"http", "silo-hybrik.pluto.tv.s3.amazonaws.com", "200 OK"},
   {"http", "siloh-fs.plutotv.net", "403 OK"},
   {"http", "siloh-ns1.plutotv.net", "403 OK"},
   {"https", "siloh-fs.plutotv.net", "403 OK"},
   {"https", "siloh-ns1.plutotv.net", "403 OK"},
}

type FileBase struct {
   Scheme string
   Host   string
   Status string
}

type Namer struct {
   Video *OnDemand
}

func (Namer) Season() int {
   return 0
}

func (Namer) Episode() int {
   return 0
}

func (Namer) Year() int {
   return 0
}

func (n Namer) Show() string {
   if v := n.Video.season; v != nil {
      return v.show.Name
   }
   return ""
}

func (n Namer) Title() string {
   return n.Video.Slug
}

type OnDemand struct {
   Episode string `json:"_id"`
   Id      string
   Name    string
   Seasons []*VideoSeason
   Slug    string
   season  *VideoSeason
}

func (o OnDemand) Clip() (*EpisodeClip, error) {
   req, err := http.NewRequest("", "https://api.pluto.tv", nil)
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

type Url struct {
   Url *url.URL
}

func (u *Url) UnmarshalText(text []byte) error {
   u.Url = &url.URL{}
   return u.Url.UnmarshalBinary(text)
}

type VideoSeason struct {
   Episodes []*OnDemand
   show     *OnDemand
}
