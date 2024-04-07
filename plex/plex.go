package plex

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type anonymous struct {
   AuthToken string
}

func (a *anonymous) New() error {
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

func (a anonymous) abs(path string, query url.Values) string {
   query.Set("x-plex-token", a.AuthToken)
   var u url.URL
   u.Host = "vod.provider.plex.tv"
   u.Path = path
   u.RawQuery = query.Encode()
   u.Scheme = "https"
   return u.String()
}

func (a anonymous) on_demand(d *discover_match) (*on_demand, error) {
   req, err := http.NewRequest("GET", "https://vod.provider.plex.tv", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/library/metadata/" + d.RatingKey
   req.Header = http.Header{
      "accept": {"application/json"},
      "x-plex-token": {a.AuthToken},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var s struct {
      MediaContainer struct {
         Metadata []on_demand
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   return &s.MediaContainer.Metadata[0], nil
}

type media_part struct {
   Key string
   License string
}

func (media_part) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (media_part) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (p media_part) RequestUrl() (string, bool) {
   return p.License, true
}

func (media_part) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

type on_demand struct {
   Media []struct {
      Part []media_part
      Protocol string
   }
}

func (o on_demand) dash(a anonymous) (*media_part, bool) {
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
