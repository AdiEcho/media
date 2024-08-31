package plex

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

type Path string

func (p Path) String() string {
   return string(p)
}

// watch.plex.tv/movie/the-hurt-locker
// https://watch.plex.tv/movie/the-hurt-locker
// https://watch.plex.tv/watch/movie/the-hurt-locker
func (p *Path) Set(s string) error {
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "watch.plex.tv")
   s = strings.TrimPrefix(s, "/watch")
   *p = Path(s)
   return nil
}

type DiscoverMatch struct {
   GrandparentTitle string
   Index int
   ParentIndex int
   RatingKey string
   Title string
   Year int
}

func (MediaPart) Error() string {
   return "MediaPart"
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

func (p *MediaPart) RequestUrl() (string, bool) {
   return p.License, true
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

type Anonymous struct {
   AuthToken string
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

///

func (o OnDemand) Dash(a Anonymous) (*MediaPart, bool) {
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

func (a Anonymous) Discover(p Path) (*DiscoverMatch, error) {
   req, err := http.NewRequest(
      "", "https://discover.provider.plex.tv/library/metadata/matches", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("accept", "application/json")
   req.URL.RawQuery = url.Values{
      "url": {p.s},
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
   var match struct {
      MediaContainer struct {
         Metadata []DiscoverMatch
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&match)
   if err != nil {
      return nil, err
   }
   return &match.MediaContainer.Metadata[0], nil
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
   var data struct {
      MediaContainer struct {
         Metadata []OnDemand
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&data)
   if err != nil {
      return nil, err
   }
   return &data.MediaContainer.Metadata[0], nil
}
