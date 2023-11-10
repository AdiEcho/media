package hulu

import (
   "bytes"
   "encoding/json"
   "net/http"
)

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
   p.Playback.DRM.Selection_Mode = "ONE"
   p.Playback.Manifest.Type = "DASH"
   p.Playback.Version = 2
   p.Playback.Segments.Selection_Mode = "ONE"
   p.Playback.Video.Codecs.Selection_Mode = "FIRST"
   p.Playback.Audio.Codecs.Values = []codec_value{
      {
         Type: "AAC",
      },
   }
   p.Playback.Video.Codecs.Values = []codec_value{
      {
         Level: "5.2",
         Profile: "HIGH",
         Type: "H264",
      },
   }
   p.Playback.DRM.Values = []drm_value{
      {
         Security_Level: "L3",
         Type: "WIDEVINE",
         Version: "MODULAR",
      },
   }
   p.Playback.Segments.Values = func() []segment_value {
      var s segment_value
      s.Encryption.Mode = "CENC"
      s.Encryption.Type = "CENC"
      s.Type = "FMP4"
      return []segment_value{s}
   }()
   body, err := json.Marshal(p)
   if err != nil {
      return nil, err
   }
   return http.Post(
      "https://play.hulu.com/v6/playlist", "application/json",
      bytes.NewReader(body),
   )
}
