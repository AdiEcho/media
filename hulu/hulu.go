package hulu

import (
   "net/http"
   "path"
   "strconv"
   "strings"
)

type Playlist struct {
   Stream_URL string
   WV_Server string
}

func (Playlist) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Playlist) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (p Playlist) RequestUrl() (string, bool) {
   return p.WV_Server, true
}

func (Playlist) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

type DeepLink struct {
   EAB_ID string
}

type ID struct {
   s string
}

func (i ID) String() string {
   return i.s
}

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
func (i *ID) Set(s string) error {
   i.s = path.Base(s)
   return nil
}

type codec_value struct {
   Height int `json:"height,omitempty"`
   Level   string `json:"level,omitempty"`
   Profile string `json:"profile,omitempty"`
   Type    string `json:"type"`
   Width int `json:"width,omitempty"`
}

type drm_value struct {
   Security_Level string `json:"security_level"`
   Type          string `json:"type"`
   Version       string `json:"version"`
}

type playlist_request struct {
   Content_EAB_ID   string `json:"content_eab_id"`
   Deejay_Device_ID int    `json:"deejay_device_id"`
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

type Details struct {
   Episode_Name string
   Episode_Number int
   Headline string
   Premiere_Date string
   Season_Number int
   Series_Name string
}

func (d Details) Show() string {
   return d.Series_Name
}

func (d Details) Season() int {
   return d.Season_Number
}

func (d Details) Episode() int {
   return d.Episode_Number
}

func (d Details) Title() string {
   if v := d.Episode_Name; v != "" {
      return v
   }
   return d.Headline
}

func (d Details) Year() int {
   if v, _, ok := strings.Cut(d.Premiere_Date, "-"); ok {
      if v, err := strconv.Atoi(v); err == nil {
         return v
      }
   }
   return 0
}
