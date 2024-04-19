package pluto

import (
   "net/http"
   "net/url"
   "encoding/json"
)

type on_demand struct {
   ID string
}

type boot_start struct {
   VOD []on_demand
}

func (b *boot_start) New(slug, forward string) error {
   req, err := http.NewRequest("GET", "https://boot.pluto.tv/v4/start", nil)
   if err != nil {
      return err
   }
   if forward != "" {
      req.Header.Set("x-forwarded-for", forward)
   }
   req.URL.RawQuery = url.Values{
      "appName": {"web"},
      "appVersion": {"9"},
      "clientID": {"9"},
      "clientModelNumber": {"9"},
      "episodeSlugs": {slug},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(b)
}
