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

func (a Authenticate) Playlist(d *DeepLink) (*Playlist, error) {
   var p playlist_request
   p.Playback.Video.Codecs.Values = []codec_value{
      {
         Height: 9999,
         Level: "9",
         Profile: "MAIN_10",
         Tier: "MAIN",
         Type: "H265",
         Width: 9999,
      },
      {
         Height: 9999,
         Level: "9",
         Profile: "HIGH",
         Type: "H264",
         Width: 9999,
      },
   }
   p.ContentEabId = d.EabId
   p.Playback.Audio.Codecs.SelectionMode = "ALL"
   p.Playback.Audio.Codecs.Values = []codec_value{
      {Type: "AAC"},
      {Type: "EC3"},
   }
   p.Playback.Drm.SelectionMode = "ALL"
   p.Playback.Drm.Values = []drm_value{
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
   p.Playback.Video.Codecs.SelectionMode = "ALL"
   p.Unencrypted = true
   p.Playback.Version = 9 // this is required for 1080p
   p.DeejayDeviceId = 166
   p.Version = 9999999
   body, err := json.MarshalIndent(p, "", " ")
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
      "authorization": {"Bearer " + a.Data.UserToken},
      "content-type": {"application/json"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   play := new(Playlist)
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

func (a Authenticate) DeepLink(id EntityId) (*DeepLink, error) {
   req, err := http.NewRequest("", "https://discover.hulu.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/content/v5/deeplink/playback"
   req.URL.RawQuery = url.Values{
      "id": {id.s},
      "namespace": {"entity"},
   }.Encode()
   req.Header.Set("authorization", "Bearer " + a.Data.UserToken)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   link := new(DeepLink)
   err = json.NewDecoder(resp.Body).Decode(link)
   if err != nil {
      return nil, err
   }
   return link, nil
}

func (a *Authenticate) New(email, password string) error {
   resp, err := http.PostForm(
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
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return errors.New(b.String())
   }
   a.raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a Authenticate) Marshal() []byte {
   return a.raw
}

type Authenticate struct {
   Data *struct {
      UserToken string `json:"user_token"`
   }
   raw []byte
}

func (a *Authenticate) Unmarshal(raw []byte) error {
   return json.Unmarshal(raw, a)
}
