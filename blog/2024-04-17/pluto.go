package pluto

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

type boot_start struct {
   Servers struct {
      StitcherDash string
   }
   SessionToken string
   VOD []struct {
      Name string
      Stitched struct {
         Paths []struct {
            Path string
            Type string
         }
      }
   }
}

func (b boot_start) mpd() (string, bool) {
   for _, vod := range b.VOD {
      for _, path := range vod.Stitched.Paths {
         if path.Type == "mpd" {
            var s strings.Builder
            s.WriteString(b.Servers.StitcherDash)
            s.WriteString("/v2")
            s.WriteString(path.Path)
            s.WriteString("?jwt=")
            s.WriteString(b.SessionToken)
            return s.String(), true
         }
      }
   }
   return "", false
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
      "drmCapabilities": {"widevine:L3"},
      "episodeSlugs": {slug},
      // if you omit these, the `sessionToken` in the response will trigger an
      // error with the MPD request
      "deviceMake": {"firefox"},
      "deviceModel": {"web"},
      "deviceVersion": {"9"},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(b)
}

var forwards = map[string]string{"Canada": "99.224.0.0"}
