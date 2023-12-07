package amc

import (
   "encoding/json"
   "errors"
   "net/http"
   "strings"
   "time"
)

func (p Playback) HTTP_DASH() (*Source, error) {
   for _, s := range p.body.Data.Playback_JSON_Data.Sources {
      if strings.HasPrefix(s.Src, "https://") {
         if s.Type == "application/dash+xml" {
            return &s, nil
         }
      }
   }
   return nil, errors.New("data.playbackJsonData.sources")
}

type Content struct {
   Data	struct {
      Children []struct {
         Properties json.RawMessage
         Type string
      }
   }
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

func (v Video) Date() (time.Time, error) {
   return time.Parse(time.RFC3339, v.Meta.Airdate)
}

func (v Video) Episode() (int64, error) {
   return v.Meta.Episode_Number, nil
}

func (v Video) Season() (int64, error) {
   return v.Meta.Season, nil
}

func (v Video) Series() string {
   return v.Meta.Show_Title
}

func (v Video) Title() string {
   return v.Text.Title
}

func (Playback) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) Response_Body(b []byte) ([]byte, error) {
   return b, nil
}

type Playback struct {
   header http.Header
   body struct {
      Data struct {
         Playback_JSON_Data struct {
            Sources []Source
         } `json:"playbackJsonData"`
      }
   }
}

func (p Playback) Request_Header() http.Header {
   return http.Header{
      "bcov-auth": {p.header.Get("X-AMCN-BC-JWT")},
   }
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

func (p Playback) Request_URL() (string, error) {
   v, err := p.HTTP_DASH()
   if err != nil {
      return "", err
   }
   return v.Key_Systems.Widevine.License_URL, nil
}

func (u URL) String() string {
   return u.path
}

// https://www.amcplus.com/movies/queen-of-earth--1026724
type URL struct {
   path string // /movies/queen-of-earth--1026724
   nid string // 1026724
}

func (u *URL) Set(s string) error {
   var found bool
   _, u.path, found = strings.Cut(s, "amcplus.com")
   if !found {
      return errors.New("amcplus.com")
   }
   _, u.nid, found = strings.Cut(s, "--")
   if !found {
      return errors.New("--")
   }
   return nil
}
