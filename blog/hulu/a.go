package hulu

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
      DRM struct {
         Selection_Mode string `json:"selection_mode"`
         Values []struct {
            Security_Level string `json:"security_level"`
            Type          string `json:"type"`
            Version       string `json:"version"`
         } `json:"values"`
      } `json:"drm"`
      Manifest struct {
         Type string `json:"type"`
      } `json:"manifest"`
      Segments struct {
         Selection_Mode string `json:"selection_mode"`
         Values []struct {
            Encryption struct {
               Mode string `json:"mode"`
               Type string `json:"type"`
            } `json:"encryption"`
            Type string `json:"type"`
         } `json:"values"`
      } `json:"segments"`
      Version int `json:"version"`
      Video   struct {
         Codecs struct {
            Selection_Mode string `json:"selection_mode"`
            Values []struct {
               Level   string `json:"level"`
               Profile string `json:"profile"`
               Type    string `json:"type"`
            } `json:"values"`
         } `json:"codecs"`
      } `json:"video"`
   } `json:"playback"`
}
