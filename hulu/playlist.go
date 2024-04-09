package hulu

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

func (a Authenticate) Playlist(d *DeepLink) (*Playlist, error) {
   var p playlist_request
   p.Content_EAB_ID = d.EAB_ID
   p.Deejay_Device_ID = 166
   p.Playback.Audio.Codecs.Selection_Mode = "ALL"
   p.Playback.Audio.Codecs.Values = []codec_value{
      {Type: "AAC"},
      {Type: "EC3"},
   }
   p.Playback.DRM.Selection_Mode = "ALL"
   p.Playback.DRM.Values = []drm_value{
      {
         Security_Level: "L3",
         Type: "WIDEVINE",
         Version: "MODULAR",
      },
   }
   p.Playback.Manifest.Type = "DASH"
   p.Playback.Segments.Selection_Mode = "ALL"
   p.Playback.Segments.Values = func() []segment_value {
      var s segment_value
      s.Encryption.Mode = "CENC"
      s.Encryption.Type = "CENC"
      s.Type = "FMP4"
      return []segment_value{s}
   }()
   p.Playback.Version = 2 // this is required for 1080p
   p.Playback.Video.Codecs.Selection_Mode = "ALL"
   p.Playback.Video.Codecs.Values = []codec_value{
      {
         Height: 9999,
         Width: 9999,
         Level: "5.2",
         Profile: "HIGH",
         Type: "H264",
      },
   }
   p.Unencrypted = true
   p.Version = 5012541
   body, err := json.Marshal(p)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://play.hulu.com/v6/playlist", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.v.Data.User_Token},
      "Content-Type": {"application/json"},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   play := new(Playlist)
   if err := json.NewDecoder(res.Body).Decode(play); err != nil {
      return nil, err
   }
   return play, nil
}

type Playlist struct {
   Stream_URL string
   WV_Server string
}

func (Playlist) RequestBody(b []byte) ([]byte, error) {
   return b, nil
}

func (Playlist) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (p Playlist) RequestUrl() (string, bool) {
   return p.WV_Server, true
}

func (Playlist) ResponseBody(b []byte) ([]byte, error) {
   return b, nil
}
