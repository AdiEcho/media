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
   "strings"
   "time"
)

func (d *DefaultToken) Login(key PublicKey, login DefaultLogin) error {
   address := func() string {
      var b bytes.Buffer
      b.WriteString("https://default.any-")
      b.WriteString(home_market)
      b.WriteString(".prd.api.discomax.com/login")
      return b.String()
   }()
   body, err := json.Marshal(login)
   if err != nil {
      return err
   }
   req, err := http.NewRequest("POST", address, bytes.NewReader(body))
   if err != nil {
      return err
   }
   req.Header.Set("authorization", "Bearer "+d.Body.Data.Attributes.Token)
   req.Header.Set("content-type", "application/json")
   req.Header.Set("x-disco-arkose-token", key.Token)
   req.Header.Set("x-disco-client-id", func() string {
      timestamp := time.Now().Unix()
      hash := hmac.New(sha256.New, default_key.Key)
      fmt.Fprintf(hash, "%v:POST:/login:%s", timestamp, body)
      signature := hash.Sum(nil)
      return fmt.Sprintf("%v:%v:%x", default_key.Id, timestamp, signature)
   }())
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
   session := make(session_state)
   session.Set(resp.Header.Get("x-wbd-session-state"))
   for key := range session {
      switch key {
      case "device", "token", "user":
      default:
         delete(session, key)
      }
   }
   d.SessionState = session.String()
   return json.NewDecoder(resp.Body).Decode(&d.Body)
}

func (d *DefaultToken) New() error {
   req, err := http.NewRequest(
      "", "https://default.any-any.prd.api.discomax.com/token?realm=bolt", nil,
   )
   if err != nil {
      return err
   }
   // fuck you Max
   req.Header.Set("x-device-info", "!/!(!/!;!/!;!)")
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
   return json.NewDecoder(resp.Body).Decode(&d.Body)
}

type DefaultToken struct {
   SessionState string
   Body struct {
      Data struct {
         Attributes struct {
            Token string
         }
      }
   }
}

func (d DefaultToken) Marshal() ([]byte, error) {
   return json.Marshal(d)
}

func (d *DefaultToken) Unmarshal(text []byte) error {
   return json.Unmarshal(text, d)
}

func (d DefaultToken) decision() (*default_decision, error) {
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
   req.Header.Set("authorization", "Bearer "+d.Body.Data.Attributes.Token)
   req.URL.Path = "/labs/api/v1/sessions/feature-flags/decisions"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   decision := new(default_decision)
   err = json.NewDecoder(resp.Body).Decode(decision)
   if err != nil {
      return nil, err
   }
   return decision, nil
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

// note you can use other keys, but you need to change home_market to match
var default_key = hmac_key{
   Id:  "android1_prd",
   Key: []byte("6fd2c4b9-7b43-49ee-a62e-57ffd7bdfe9c"),
}

type session_state map[string]string

func (s session_state) Set(text string) error {
   for text != "" {
      var key string
      key, text, _ = strings.Cut(text, ";")
      key, value, _ := strings.Cut(key, ":")
      s[key] = value
   }
   return nil
}

func (s session_state) String() string {
   var (
      b strings.Builder
      sep bool
   )
   for key, value := range s {
      if sep {
         b.WriteByte(';')
      } else {
         sep = true
      }
      b.WriteString(key)
      b.WriteByte(':')
      b.WriteString(value)
   }
   return b.String()
}

type hmac_key struct {
   Id  string
   Key []byte
}

type DefaultLogin struct {
   Credentials struct {
      Username string `json:"username"`
      Password string `json:"password"`
   } `json:"credentials"`
}

const home_market = "amer"

type default_decision struct {
   HmacKeys struct {
      Config struct {
         Android   *hmac_key
         AndroidTv *hmac_key
         FireTv    *hmac_key
         Hwa       *hmac_key
         Ios       *hmac_key
         TvOs      *hmac_key
         Web       *hmac_key
      }
   }
}

const arkose_site_key = "B0217B00-2CA4-41CC-925D-1EEB57BFFC2F"

func (d DefaultRoutes) video() (*route_include, bool) {
   for _, include := range d.Included {
      if include.Id == d.Data.Attributes.Url.VideoId {
         return &include, true
      }
   }
   return nil, false
}

func (d DefaultRoutes) Season() int {
   if v, ok := d.video(); ok {
      return v.Attributes.SeasonNumber
   }
   return 0
}

func (d DefaultRoutes) Episode() int {
   if v, ok := d.video(); ok {
      return v.Attributes.EpisodeNumber
   }
   return 0
}

func (d DefaultRoutes) Title() string {
   if v, ok := d.video(); ok {
      return v.Attributes.Name
   }
   return ""
}

func (d DefaultRoutes) Year() int {
   if v, ok := d.video(); ok {
      return v.Attributes.AirDate.Year()
   }
   return 0
}

func (d DefaultRoutes) Show() string {
   if v, ok := d.video(); ok {
      if v.Attributes.SeasonNumber >= 1 {
         for _, include := range d.Included {
            if include.Id == v.Relationships.Show.Data.Id {
               return include.Attributes.Name
            }
         }
      }
   }
   return ""
}

type PublicKey struct {
   Token string
}

func (p *PublicKey) New() error {
   resp, err := http.PostForm(
      "https://wbd-api.arkoselabs.com/fc/gt2/public_key/"+arkose_site_key,
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

type route_include struct {
   Attributes struct {
      AirDate       time.Time
      Name          string
      EpisodeNumber int
      SeasonNumber  int
   }
   Id            string
   Relationships *struct {
      Show *struct {
         Data struct {
            Id string
         }
      }
   }
}
