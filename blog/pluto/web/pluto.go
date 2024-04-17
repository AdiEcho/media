package pluto

import (
   "encoding/json"
   "net/http"
   "net/url"
)

var forwards = map[string]string{"Canada": "99.224.0.0"}

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
      "drmCapabilities": {"widevine:L3"},
      "episodeSlugs": {slug},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(b)
}

type boot_start struct {
   Servers struct {
      StitcherDash string
   }
   VOD []struct {
      Name string
      Stitched struct {
         Paths []struct {
            Path string
         }
      }
   }
}
