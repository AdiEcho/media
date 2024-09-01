package plex

import (
   "net/http"
   "net/url"
   "strings"
)

type DiscoverMatch struct {
   GrandparentTitle string
   Index int
   ParentIndex int
   RatingKey string
   Title string
   Year int
}

type Namer struct {
   Match *DiscoverMatch
}

func (n *Namer) Episode() int {
   return n.Match.Index
}

func (n *Namer) Season() int {
   return n.Match.ParentIndex
}

func (n *Namer) Show() string {
   return n.Match.GrandparentTitle
}

func (n *Namer) Title() string {
   return n.Match.Title
}

func (n *Namer) Year() int {
   return n.Match.Year
}

type Address struct {
   Path string
}

func (a Address) String() string {
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

func (MediaPart) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (MediaPart) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (MediaPart) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
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

func (u *Url) UnmarshalText(text []byte) error {
   u.Url = &url.URL{}
   err := u.Url.UnmarshalBinary(text)
   if err != nil {
      return err
   }
   u.Url.Scheme = "https"
   u.Url.Host = "vod.provider.plex.tv"
   return nil
}

type MediaPart struct {
   Key Url
   License *Url
}

func (m *MediaPart) RequestUrl() (string, bool) {
   return m.License.Url.String(), true
}
