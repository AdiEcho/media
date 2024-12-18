package plex

import (
   "bytes"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func (m *MediaPart) Wrap(data []byte) ([]byte, error) {
   var req http.Request
   req.Body = io.NopCloser(bytes.NewReader(data))
   req.Method = "POST"
   req.URL = m.License.Url
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

type MediaPart struct {
   Key Url
   License *Url
}

type Address struct {
   Path string
}

func (a *Address) String() string {
   return a.Path
}

func (a *Address) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "watch.plex.tv")
   a.Path = strings.TrimPrefix(s, "/watch")
   return nil
}

func (o *OnDemand) Dash() (*MediaPart, bool) {
   for _, media := range o.Media {
      if media.Protocol == "dash" {
         for _, part := range media.Part {
            return &part, true
         }
      }
   }
   return nil, false
}

type OnDemand struct {
   Media []struct {
      Part []MediaPart
      Protocol string
   }
}

type Url struct {
   Url *url.URL
}

func (u *Url) UnmarshalText(data []byte) error {
   u.Url = &url.URL{}
   err := u.Url.UnmarshalBinary(data)
   if err != nil {
      return err
   }
   u.Url.Scheme = "https"
   u.Url.Host = "vod.provider.plex.tv"
   return nil
}

func (n Namer) Show() string {
   return n.Match.GrandparentTitle
}

func (n Namer) Title() string {
   return n.Match.Title
}

type Namer struct {
   Match *DiscoverMatch
}

type DiscoverMatch struct {
   GrandparentTitle string
   Index int
   ParentIndex int
   RatingKey string
   Title string
   Year int
}

func (n Namer) Episode() int {
   return n.Match.Index
}

func (n Namer) Season() int {
   return n.Match.ParentIndex
}

func (n Namer) Year() int {
   return n.Match.Year
}
