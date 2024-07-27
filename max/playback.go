package max

import (
   "bytes"
   "encoding/json"
   "errors"
   "net/http"
)

func (d DefaultToken) Playback(web WebAddress) (*Playback, error) {
   body, err := func() ([]byte, error) {
      var p playback_request
      p.ConsumptionType = "streaming"
      p.EditId = web.EditId
      return json.MarshalIndent(p, "", " ")
   }()
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://default.any-any.prd.api.discomax.com",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      var b bytes.Buffer
      b.WriteString("/playback-orchestrator/any/playback-orchestrator/v1")
      b.WriteString("/playbackInfo")
      return b.String()
   }()
   req.Header = http.Header{
      "authorization": {"Bearer "+d.Body.Data.Attributes.Token},
      "content-type": {"application/json"},
      "x-wbd-session-state": {d.SessionState},
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
   play := new(Playback)
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

type Playback struct {
   Drm struct {
      Schemes struct {
         Widevine struct {
            LicenseUrl string
         }
      }
   }
   Manifest struct {
      Url Url
   }
}

func (Playback) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (p Playback) RequestUrl() (string, bool) {
   return p.Drm.Schemes.Widevine.LicenseUrl, true
}

type playback_request struct {
   AppBundle            string `json:"appBundle"`            // required
   ApplicationSessionId string `json:"applicationSessionId"` // required
   Capabilities         struct {
      Manifests struct {
         Formats struct {
            Dash struct{} `json:"dash"` // required
         } `json:"formats"` // required
      } `json:"manifests"` // required
   } `json:"capabilities"` // required
   ConsumptionType string `json:"consumptionType"`
   DeviceInfo      struct {
      Player struct {
         MediaEngine struct {
            Name    string `json:"name"`    // required
            Version string `json:"version"` // required
         } `json:"mediaEngine"` // required
         PlayerView struct {
            Height int `json:"height"` // required
            Width  int `json:"width"`  // required
         } `json:"playerView"` // required
         Sdk struct {
            Name    string `json:"name"`    // required
            Version string `json:"version"` // required
         } `json:"sdk"` // required
      } `json:"player"` // required
   } `json:"deviceInfo"` // required
   EditId            string   `json:"editId"`
   FirstPlay         bool     `json:"firstPlay"`         // required
   Gdpr              bool     `json:"gdpr"`              // required
   PlaybackSessionId string   `json:"playbackSessionId"` // required
   UserPreferences   struct{} `json:"userPreferences"`   // required
}
