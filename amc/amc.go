package amc

import (
   "encoding/json"
   "errors"
   "net/http"
   "strings"
   "time"
)

func (p Path) nid() (string, error) {
   _, nid, found := strings.Cut(p.s, "--")
   if !found {
      return "", errors.New("nid")
   }
   return nid, nil
}

type Path struct {
   s string
}

func (p Path) String() string {
   return p.s
}

func (p *Path) Set(s string) error {
   if _, after, found := strings.Cut(s, "://"); found {
      s = after // remove scheme
   }
   if i := strings.IndexByte(s, '/'); i >= 1 {
      s = s[i:] // remove host
   }
   p.s = s
   return nil
}

func (c Content) Video() (*Video, error) {
   for _, child := range c.Data.Children {
      if child.Type == "video-player-ap" {
         var s struct {
            Current_Video Video `json:"currentVideo"`
         }
         err := json.Unmarshal(child.Properties, &s)
         if err != nil {
            return nil, err
         }
         return &s.Current_Video, nil
      }
   }
   return nil, errors.New("video-player-ap")
}

type Video struct {
   Meta struct {
      Show_Title string `json:"showTitle"`
      Season int64 `json:",string"`
      Episode_Number int64 `json:"episodeNumber"`
      Airdate string // 1996-01-01T00:00:00.000Z
   }
   Text struct {
      Title string
   }
}

func (v Video) Series() string {
   return v.Meta.Show_Title
}

func (v Video) Season() (int64, error) {
   return v.Meta.Season, nil
}

func (v Video) Episode() (int64, error) {
   return v.Meta.Episode_Number, nil
}

func (v Video) Title() string {
   return v.Text.Title
}

func (v Video) Date() (time.Time, error) {
   return time.Parse(time.RFC3339, v.Meta.Airdate)
}

type Playback struct {
   h http.Header
   sources []Source
}

func (p Playback) HTTP_DASH() *Source {
   for _, source := range p.sources {
      if strings.HasPrefix(source.Src, "http://") {
         if source.Type == "application/dash+xml" {
            return &source
         }
      }
   }
   return nil
}

type Source struct {
   Key_Systems *struct {
      Widevine struct {
         License_URL string
      } `json:"com.widevine.alpha"`
   }
   Src string
   Type string
}

type Content struct {
   Data	struct {
      Children []struct {
         Properties json.RawMessage
         Type string
      }
   }
}

func (p Playback) Request_URL() string {
   return p.HTTP_DASH().Key_Systems.Widevine.License_URL
}

func (Playback) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (p Playback) Request_Header() http.Header {
   return http.Header{
      "bcov-auth": {p.h.Get("X-AMCN-BC-JWT")},
   }
}
