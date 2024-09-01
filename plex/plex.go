package plex

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

type Anonymous struct {
   AuthToken string
}

func (a *Anonymous) New() error {
   req, err := http.NewRequest(
      "POST", "https://plex.tv/api/v2/users/anonymous", nil,
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "accept": {"application/json"},
      "x-plex-product": {"Plex Mediaverse"},
      "x-plex-client-identifier": {"!"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(a)
}

func (a Anonymous) abs(path string, query url.Values) string {
   query.Set("x-plex-token", a.AuthToken)
   var u url.URL
   u.Host = "vod.provider.plex.tv"
   u.Path = path
   u.RawQuery = query.Encode()
   u.Scheme = "https"
   return u.String()
}

func (a Anonymous) Discover(u Url) (*DiscoverMatch, error) {
   req, err := http.NewRequest(
      "", "https://discover.provider.plex.tv/library/metadata/matches", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("accept", "application/json")
   req.URL.RawQuery = url.Values{
      "url": {u.Path},
      "x-plex-token": {a.AuthToken},
   }.Encode()
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   var value struct {
      MediaContainer struct {
         Metadata []DiscoverMatch
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.MediaContainer.Metadata[0], nil
}

func (a Anonymous) Video(
   match *DiscoverMatch, forward string,
) (*OnDemand, error) {
   req, err := http.NewRequest("", "https://vod.provider.plex.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/library/metadata/" + match.RatingKey
   req.Header.Set("accept", "application/json")
   req.Header.Set("x-plex-token", a.AuthToken)
   if forward != "" {
      req.Header.Set("x-forwarded-for", forward)
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      return nil, errors.New(resp.Status)
   }
   var value struct {
      MediaContainer struct {
         Metadata []OnDemand
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&value)
   if err != nil {
      return nil, err
   }
   return &value.MediaContainer.Metadata[0], nil
}
type OnDemand struct {
   Media []struct {
      Part []MediaPart
      Protocol string
   }
}

type MediaPart struct {
   Key string
   License string
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
