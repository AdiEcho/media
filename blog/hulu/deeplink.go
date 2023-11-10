package hulu

import (
   "encoding/json"
   "net/http"
   "net/url"
)

type deep_link struct {
   EAB_ID string
}

func (a authenticate) deeplink(id string) (*deep_link, error) {
   req, err := http.NewRequest("GET", "https://discover.hulu.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/content/v5/deeplink/playback"
   req.URL.RawQuery = url.Values{
      "id": {id},
      "namespace": {"entity"},
      "user_token": {a.Data.User_Token},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   link := new(deep_link)
   if err := json.NewDecoder(res.Body).Decode(link); err != nil {
      return nil, err
   }
   return link, nil
}

var playlist_body = strings.NewReader(`
{
   "content_eab_id": "EAB::023c49bf-6a99-4c67-851c-4c9e7609cc1d::196861183::262714326",
   "deejay_device_id": 166,
   "token": "DhA8XQLI1A9DkskfH3IaZ7L/aRU-nlLJF0thPd3yFttJT6_Ivg--uEEzNxGQrnh...",
   "unencrypted": true,
   "version": 5012541,
   "playback": {
      "audio": {
         "codecs": {
            "selection_mode": "ONE",
            "values": [
               {
                  "type": "AAC"
               }
            ]
         }
      },
      "drm": {
         "selection_mode": "ONE",
         "values": [
            {
               "security_level": "L3",
               "type": "WIDEVINE",
               "version": "MODULAR"
            }
         ]
      },
      "manifest": {
         "type": "DASH"
      },
      "segments": {
         "selection_mode": "ONE",
         "values": [
            {
               "encryption": {
                  "mode": "CENC",
                  "type": "CENC"
               },
               "type": "FMP4"
            }
         ]
      },
      "version": 2,
      "video": {
         "codecs": {
            "selection_mode": "FIRST",
            "values": [
               {
                  "level": "5.2",
                  "profile": "HIGH",
                  "type": "H264"
               }
            ]
         }
      }
   }
}
`)
