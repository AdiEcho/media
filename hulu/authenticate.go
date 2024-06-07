package hulu

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
   "strings"
)

type Authenticate struct {
   Data []byte
   v struct {
      Data struct {
         UserToken string `json:"user_token"`
      }
   }
}

func (a *Authenticate) Unmarshal() error {
   return json.Unmarshal(a.Data, &a.v)
}

func (a Authenticate) DeepLink(watch ID) (*DeepLink, error) {
   req, err := http.NewRequest("GET", "https://discover.hulu.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/content/v5/deeplink/playback"
   req.URL.RawQuery = url.Values{
      "id": {watch.s},
      "namespace": {"entity"},
   }.Encode()
   req.Header.Set("Authorization", "Bearer " + a.v.Data.UserToken)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   link := new(DeepLink)
   err = json.NewDecoder(res.Body).Decode(link)
   if err != nil {
      return nil, err
   }
   return link, nil
}
func (a *Authenticate) New(email, password string) error {
   res, err := http.PostForm(
      "https://auth.hulu.com/v2/livingroom/password/authenticate", url.Values{
         "friendly_name": {"!"},
         "password": {password},
         "serial_number": {"!"},
         "user_email": {email},
      },
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return errors.New(b.String())
   }
   a.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}
func (a Authenticate) Playlist(d *DeepLink) (*Playlist, error) {
   var p playlist_request
   p.ContentEabId = d.EabId
   p.DeejayDeviceId = 166
   p.Playback.Audio.Codecs.SelectionMode = "ALL"
   p.Playback.Audio.Codecs.Values = []codec_value{
      {Type: "AAC"},
      {Type: "EC3"},
   }
   p.Playback.DRM.SelectionMode = "ALL"
   p.Playback.DRM.Values = []drm_value{
      {
         SecurityLevel: "L3",
         Type: "WIDEVINE",
         Version: "MODULAR",
      },
   }
   p.Playback.Manifest.Type = "DASH"
   p.Playback.Segments.SelectionMode = "ALL"
   p.Playback.Segments.Values = func() []segment_value {
      var s segment_value
      s.Encryption.Mode = "CENC"
      s.Encryption.Type = "CENC"
      s.Type = "FMP4"
      return []segment_value{s}
   }()
   p.Playback.Version = 2 // this is required for 1080p
   p.Playback.Video.Codecs.SelectionMode = "ALL"
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
      "Authorization": {"Bearer " + a.v.Data.UserToken},
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
   err = json.NewDecoder(res.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

