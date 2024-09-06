package amc

import (
   "encoding/base64"
   "errors"
   "net/http"
   "strings"
   "time"
)

func cache_hash() string {
   return base64.StdEncoding.EncodeToString([]byte("ff="))
}

type Address struct {
   Nid string
   Path string
}

func (a *Address) String() string {
   return a.Path
}

func (a *Address) Set(text string) error {
   var found bool
   _, a.Path, found = strings.Cut(text, "amcplus.com")
   if !found {
      return errors.New("amcplus.com")
   }
   _, a.Nid, found = strings.Cut(a.Path, "--")
   if !found {
      return errors.New("--")
   }
   return nil
}

func (ContentCompiler) VideoError() error {
   return errors.New("ContentCompiler.Video")
}

type ContentCompiler struct {
   Children []struct {
      Properties struct {
         CurrentVideo CurrentVideo
      }
      Type string
   }
}

func (c *ContentCompiler) Video() (*CurrentVideo, bool) {
   for _, child := range c.Children {
      if child.Type == "video-player-ap" {
         return &child.Properties.CurrentVideo, true
      }
   }
   return nil, false
}

type CurrentVideo struct {
   Meta struct {
      Airdate time.Time // 1996-01-01T00:00:00.000Z
      EpisodeNumber int
      Season int `json:",string"`
      ShowTitle string
   }
   Text struct {
      Title string
   }
}

func (c *CurrentVideo) Title() string {
   return c.Text.Title
}

func (c *CurrentVideo) Year() int {
   return c.Meta.Airdate.Year()
}

func (c *CurrentVideo) Episode() int {
   return c.Meta.EpisodeNumber
}

func (c *CurrentVideo) Show() string {
   return c.Meta.ShowTitle
}

func (c *CurrentVideo) Season() int {
   return c.Meta.Season
}

type DataSource struct {
   KeySystems *struct {
      Widevine struct {
         LicenseUrl string `json:"license_url"`
      } `json:"com.widevine.alpha"`
   } `json:"key_systems"`
   Src string
   Type string
}

func (Playback) DashError() error {
   return errors.New("Playback.DashError")
}

func (Playback) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (p *Playback) RequestUrl() (string, bool) {
   if v, ok := p.Dash(); ok {
      return v.KeySystems.Widevine.LicenseUrl, true
   }
   return "", false
}

func (Playback) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (p *Playback) RequestHeader() (http.Header, error) {
   head := http.Header{}
   head.Set("bcov-auth", p.AmcnBcJwt)
   return head, nil
}

func (p *Playback) Dash() (*DataSource, bool) {
   for _, source := range p.Data.PlaybackJsonData.Sources {
      if source.Type == "application/dash+xml" {
         return &source, true
      }
   }
   return nil, false
}

type Playback struct {
   AmcnBcJwt string `json:"-"`
   Data struct {
      PlaybackJsonData struct {
         Sources []DataSource
      }
   }
}
