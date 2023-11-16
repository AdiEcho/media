package hulu

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "path"
)

func (a Authenticate) Details(d *Deep_Link) (*Details, error) {
   body, err := func() ([]byte, error) {
      m := map[string][]string{
         "eabs": {d.EAB_ID},
      }
      return json.Marshal(m)
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://guide.hulu.com/guide/details", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = "user_token=" + a.Value.Data.User_Token
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   var s struct {
      Items []Details
   }
   if err := json.NewDecoder(res.Body).Decode(&s); err != nil {
      return nil, err
   }
   if len(s.Items) == 0 {
      return nil, errors.New("items length is zero")
   }
   return &s.Items[0], nil
}

func (d Details) Series() string {
   return d.Series_Name
}

func (d Details) Title() string {
   return d.Episode_Name
}

type Details struct {
   Episode_Name string
   Episode_Number int64
   Season_Number int64
   Series_Name string
}

func (d Details) Season() (int64, error) {
   return d.Season_Number, nil
}

func (d Details) Episode() (int64, error) {
   return d.Episode_Number, nil
}

func (a Authenticate) Playlist(d *Deep_Link) (*Playlist, error) {
   var p playlist_request
   p.Content_EAB_ID = d.EAB_ID
   p.Deejay_Device_ID = 166
   p.Token = a.Value.Data.User_Token
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
   res, err := http.Post(
      "https://play.hulu.com/v6/playlist", "application/json",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   play := new(Playlist)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
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

type Playlist struct {
   Stream_URL string
   WV_Server string
}

func (Playlist) Request_Body(b []byte) ([]byte, error) {
   return b, nil
}

func (Playlist) Request_Header() http.Header {
   return nil
}

func (p Playlist) Request_URL() string {
   return p.WV_Server
}

func (Playlist) Response_Body(b []byte) ([]byte, error) {
   return b, nil
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

type Deep_Link struct {
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

func (a Authenticate) Deep_Link(watch ID) (*Deep_Link, error) {
   req, err := http.NewRequest("GET", "https://discover.hulu.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/content/v5/deeplink/playback"
   req.URL.RawQuery = url.Values{
      "id": {watch.s},
      "namespace": {"entity"},
      "user_token": {a.Value.Data.User_Token},
   }.Encode()
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   link := new(Deep_Link)
   if err := json.NewDecoder(res.Body).Decode(link); err != nil {
      return nil, err
   }
   return link, nil
}
