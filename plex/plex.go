package plex

import (
   "net/http"
   "net/url"
   "strings"
)

func (o *OnDemand) Dash(a Anonymous) (*MediaPart, bool) {
   for _, media := range o.Media {
      if media.Protocol == "dash" {
         p := media.Part[0]
         p.Key = a.abs(p.Key, url.Values{})
         p.License = a.abs(p.License, url.Values{
            "x-plex-drm": {"widevine"},
         })
         return &p, true
      }
   }
   return nil, false
}

type DiscoverMatch struct {
   GrandparentTitle string
   Index int
   ParentIndex int
   RatingKey string
   Title string
   Year int
}

type MediaPart struct {
   Key string
   License string
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

func (MediaPart) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (MediaPart) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (m *MediaPart) RequestUrl() (string, bool) {
   return m.License, true
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

func (u Url) String() string {
   return u.Path
}

// watch.plex.tv/movie/the-hurt-locker
// https://watch.plex.tv/movie/the-hurt-locker
// https://watch.plex.tv/watch/movie/the-hurt-locker
func (u *Url) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "watch.plex.tv")
   u.Path = strings.TrimPrefix(s, "/watch")
   return nil
}

type Url struct {
   Path string
}
