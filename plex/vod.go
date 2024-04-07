package plex

import (
   "encoding/json"
   "net/http"
   "net/url"
)

func (m vod) dash(a anonymous) (*part, bool) {
   for _, media := range m.Media {
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

type vod struct {
   Media []struct {
      Part []part
      Protocol string
   }
}

func (a anonymous) vod(d *discover) (*vod, error) {
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
         Metadata []vod
      }
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   return &s.MediaContainer.Metadata[0], nil
}

func (part) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (part) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (part) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}

type part struct {
   Key string
   License string
}

func (p part) RequestUrl() (string, bool) {
   return p.License, true
}

