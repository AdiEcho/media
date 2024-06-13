package max

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/json"
   "errors"
   "fmt"
   "net/http"
   "net/url"
   "time"
)

func (d default_token) playback(p playback_request) (*http.Response, error) {
   body, err := json.Marshal(p)
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
   req.Header.Set("content-type", "application/json")
   req.URL.Path = func() string {
      var b bytes.Buffer
      b.WriteString("/playback-orchestrator/any/playback-orchestrator/v1")
      b.WriteString("/playbackInfo")
      return b.String()
   }()
   req.Header.Set("authorization", "Bearer " + d.Data.Attributes.Token)
   return http.DefaultClient.Do(req)
}

func (p *playback_request) New() {
   p.ConsumptionType = "streaming"
   p.EditId = "1623fe4c-ef6e-4dd1-a10c-4a181f5f6579"
}

type playback_request struct {
   AppBundle string `json:"appBundle"` // required
   ApplicationSessionId string `json:"applicationSessionId"` // required
   Capabilities struct {
      Manifests struct {
         Formats struct {
            Dash struct{} `json:"dash"` // required
         } `json:"formats"` // required
      } `json:"manifests"` // required
   } `json:"capabilities"` // required
   ConsumptionType string `json:"consumptionType"`
   DeviceInfo struct {
      Player struct {
         MediaEngine struct {
            Name string `json:"name"` // required
            Version string `json:"version"` // required
         } `json:"mediaEngine"` // required
         PlayerView struct {
            Height int `json:"height"` // required
            Width int `json:"width"` // required
         } `json:"playerView"` // required
         Sdk struct {
            Name string `json:"name"` // required
            Version string `json:"version"` // required
         } `json:"sdk"` // required
      } `json:"player"` // required
   } `json:"deviceInfo"` // required
   EditId string `json:"editId"`
   FirstPlay bool `json:"firstPlay"` // required
   Gdpr bool `json:"gdpr"` // required
   PlaybackSessionId string `json:"playbackSessionId"` // required
   UserPreferences struct{} `json:"userPreferences"` // required
}

type default_login struct {
   Credentials struct {
      Username string `json:"username"`
      Password string `json:"password"`
   } `json:"credentials"`
}

type public_key struct {
   Token string
}

const arkose_site_key = "B0217B00-2CA4-41CC-925D-1EEB57BFFC2F"

func (d *default_token) unmarshal(text []byte) error {
   return json.Unmarshal(text, d)
}

func (d default_token) marshal() ([]byte, error) {
   return json.MarshalIndent(d, "", " ")
}

type default_token struct {
   Data struct {
      Attributes struct {
         Token string
      }
   }
}

type key_config struct {
   Id string
   Key []byte
}

var android_config = key_config{
   Id: "android1_prd",
   Key: []byte("6fd2c4b9-7b43-49ee-a62e-57ffd7bdfe9c"),
}

func (p *public_key) New() error {
   resp, err := http.PostForm(
      "https://wbd-api.arkoselabs.com/fc/gt2/public_key/" + arkose_site_key,
      url.Values{
         "public_key": {arkose_site_key},
      },
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(p)
}

func (d *default_token) New() error {
   req, err := http.NewRequest(
      "", "https://default.any-any.prd.api.discomax.com/token?realm=bolt", nil,
   )
   if err != nil {
      return err
   }
   req.Header.Set(
      "x-device-info",
      "BEAM-Android/4.1.1 (unknown/Android SDK built for x86; ANDROID/6.0; bafeab496eee63aa/b6746ddc-7bc7-471f-a16c-f6aaf0c34d26)",
   )
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b bytes.Buffer
      resp.Write(&b)
      return errors.New(b.String())
   }
   return json.NewDecoder(resp.Body).Decode(d)
}

func (d default_token) config() (*key_config, error) {
   body, err := json.Marshal(map[string]string{
      "projectId": "d8665e86-8706-415d-8d84-d55ceddccfb5",
   })
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
   req.Header.Set("authorization", "Bearer " + d.Data.Attributes.Token)
   req.URL.Path = "/labs/api/v1/sessions/feature-flags/decisions"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var decision struct {
      HmacKeys struct {
         Config struct {
            Android key_config
         }
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&decision)
   if err != nil {
      return nil, err
   }
   return &decision.HmacKeys.Config.Android, nil
}

func (d *default_token) login(key public_key, login default_login) error {
   body, err := json.Marshal(login)
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://default.any-amer.prd.api.discomax.com/login",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header.Set("content-type", "application/json")
   req.Header.Set("x-disco-arkose-token", key.Token)
   req.Header.Set("authorization", "Bearer " + d.Data.Attributes.Token)
   req.Header.Set("x-disco-client-id", func() string {
      timestamp := time.Now().Unix()
      hash := hmac.New(sha256.New, android_config.Key)
      fmt.Fprintf(hash, "%v:POST:/login:%s", timestamp, body)
      signature := hash.Sum(nil)
      return fmt.Sprintf("%v:%v:%x", android_config.Id, timestamp, signature)
   }())
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(d)
}
