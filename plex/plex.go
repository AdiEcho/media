package plex

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
)

func (a Anonymous) Video(d *DiscoverMatch, forward string) (*OnDemand, error) {
   req, err := http.NewRequest("GET", "https://vod.provider.plex.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/library/metadata/" + d.RatingKey
   req.Header.Set("accept", "application/json")
   req.Header.Set("x-plex-token", a.AuthToken)
   if forward != "" {
      req.Header.Set("x-forwarded-for", forward)
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var s struct {
      MediaContainer struct {
         Metadata []OnDemand
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   return &s.MediaContainer.Metadata[0], nil
}

type MediaPart struct {
   Key string
   License string
}

func (MediaPart) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (MediaPart) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (p MediaPart) RequestUrl() (string, bool) {
   return p.License, true
}

func (MediaPart) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

type OnDemand struct {
   Media []struct {
      Part []MediaPart
      Protocol string
   }
}

func (o OnDemand) DASH(a Anonymous) (*MediaPart, bool) {
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
      "Accept": {"application/json"},
      "X-Plex-Product": {"Plex Mediaverse"},
      "X-Plex-Client-Identifier": {"!"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
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
