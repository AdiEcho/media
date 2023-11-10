package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "play.hulu.com"
   req.URL.Path = "/v6/playlist"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
   req.Header["Content-Type"] = []string{"application/json"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}

var req_body = strings.NewReader(`
{
   "token": "DhA8XQLI1A9DkskfH3IaZ7L/aRU-nlLJF0thPd3yFttJT6_Ivg--uEEzNxGQrnh...",
   "version": 5012541,
   "deejay_device_id": 166,
   "content_eab_id": "EAB::023c49bf-6a99-4c67-851c-4c9e7609cc1d::196861183::262714326",
   "unencrypted": true,
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
