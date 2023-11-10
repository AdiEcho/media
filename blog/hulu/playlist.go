package hulu

import (
   "io"
   "net/http"
   "net/url"
   "strings"
)

const playlist_body = `
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
      "version": 2
   }
}
`

type playlist_request struct {
   Content_EAB_ID   string `json:"content_eab_id"`
   Deejay_Device_ID int    `json:"deejay_device_id"`
   Token          string `json:"token"`
   Unencrypted    bool   `json:"unencrypted"`
   Version        int    `json:"version"`
   Playback       struct {
      Audio struct {
         Codecs struct {
            Selection_Mode string `json:"selection_mode"`
            Values []codec_value `json:"values"`
         } `json:"codecs"`
      } `json:"audio"`
      Video   struct {
         Codecs struct {
            Selection_Mode string `json:"selection_mode"`
            Values []codec_value `json:"values"`
         } `json:"codecs"`
      } `json:"video"`
      DRM struct {
         Selection_Mode string `json:"selection_mode"`
         Values []drm_value `json:"values"`
      } `json:"drm"`
      Manifest struct {
         Type string `json:"type"`
      } `json:"manifest"`
      Segments struct {
         Selection_Mode string `json:"selection_mode"`
         Values []segment_value `json:"values"`
      } `json:"segments"`
      Version int `json:"version"`
   } `json:"playback"`
}

type codec_value struct {
   Level   string `json:"level"`
   Profile string `json:"profile"`
   Type    string `json:"type"`
}

type drm_value struct {
   Security_Level string `json:"security_level"`
   Type          string `json:"type"`
   Version       string `json:"version"`
}

type segment_value struct {
   Encryption struct {
      Mode string `json:"mode"`
      Type string `json:"type"`
   } `json:"encryption"`
   Type string `json:"type"`
}

func (a authenticate) playlist(d deep_link) (*http.Response, error) {
   var p playlist_request
   p.Content_EAB_ID = d.EAB_ID
   p.Deejay_Device_ID = 166
   p.Token = a.Data.User_Token
   p.Unencrypted = true
   p.Version = 5012541
   p.Playback.Audio.Codecs.Selection_Mode = "ONE"
   //p.playback.audio.codecs.values[0].type = "AAC";
   //p.playback.drm.selection_mode = "ONE";
   //p.playback.drm.values[0].security_level = "L3";
   //p.playback.drm.values[0].type = "WIDEVINE";
   //p.playback.drm.values[0].version = "MODULAR";
   //p.playback.manifest.type = "DASH";
   //p.playback.segments.selection_mode = "ONE";
   //p.playback.segments.values[0].encryption.mode = "CENC";
   //p.playback.segments.values[0].encryption.type = "CENC";
   //p.playback.segments.values[0].type = "FMP4";
   //p.playback.version = 2;
   //p.playback.video.codecs.selection_mode = "FIRST";
   //p.playback.video.codecs.values[0].level = "5.2";
   //p.playback.video.codecs.values[0].profile = "HIGH";
   //p.playback.video.codecs.values[0].type = "H264";
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Scheme = "https"
   req.URL.Host = "play.hulu.com"
   req.URL.Path = "/v6/playlist"
   req.Body = io.NopCloser(strings.NewReader(playlist_body))
   req.Header["Content-Type"] = []string{"application/json"}
   return new(http.Transport).RoundTrip(&req)
}
