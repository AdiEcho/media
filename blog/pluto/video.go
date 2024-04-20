package pluto

import (
   "encoding/json"
   "fmt"
   "net/http"
   "net/url"
)

var forwards = map[string]string{"Canada": "99.224.0.0"}

type on_demand struct {
   ID string
   Slug string
}

func new_video(slug, forward string) (*on_demand, error) {
   req, err := http.NewRequest("GET", "https://boot.pluto.tv/v4/start", nil)
   if err != nil {
      return nil, err
   }
   if forward != "" {
      req.Header.Set("x-forwarded-for", forward)
   }
   req.URL.RawQuery = url.Values{
      "appName": {"web"},
      "appVersion": {"9"},
      "clientID": {"9"},
      "clientModelNumber": {"9"},
      "drmCapabilities": {"widevine:L3"},
      "episodeSlugs": {slug},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var s struct {
      VOD []on_demand
   }
   err = json.NewDecoder(res.Body).Decode(&s)
   if err != nil {
      return nil, err
   }
   video := s.VOD[0]
   if video.ID != slug {
      if video.Slug != slug {
         return nil, fmt.Errorf("%+v", video)
      }
   }
   return &video, nil
}
