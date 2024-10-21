package max

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/base64"
   "encoding/json"
   "errors"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
   "time"
)

func (d *DefaultToken) Login(key PublicKey, login DefaultLogin) error {
   data, err := json.Marshal(login)
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", home_market + "/login", bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   req.Header.Set("authorization", "Bearer " + d.Token.Value)
   req.Header.Set("content-type", "application/json")
   req.Header.Set("user-agent", "BEAM-Android/4.12.0 (Google/Android SDK built for x86)")
   req.Header.Set("x-device-info", "BEAM-Android/4.12.0 (Google/Android SDK built for x86; ANDROID/7.0; ab3d8214d8d3f368/b6746ddc-7bc7-471f-a16c-f6aaf0c34d26)")
   req.Header.Set("x-disco-arkose-sitekey", "B0217B00-2CA4-41CC-925D-1EEB57BFFC2F")
   req.Header.Set("x-disco-arkose-token", key.Token)
   req.Header.Set("x-disco-client", "ANDROID:7.0:beam:4.12.0")
   req.Header.Set("x-disco-client-id", func() string {
      timestamp := time.Now().Unix()
      hash := hmac.New(sha256.New, default_key.Key)
      fmt.Fprintf(hash, "%v:POST:/login:%s", timestamp, data)
      signature := hash.Sum(nil)
      return fmt.Sprintf("%v:%v:%x", default_key.Id, timestamp, signature)
   }())
   req.Header.Set("x-disco-params", "realm=bolt,bid=beam,features=ar,rr")
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
   d.Token.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   d.Session.Raw = []byte(resp.Header.Get("x-wbd-session-state"))
   return nil
}

func (d *DefaultToken) Unmarshal() error {
   d.Session.Value = SessionState{}
   d.Session.Value.Set(string(d.Session.Raw))
   var data struct {
      Data struct {
         Attributes struct {
            Token string
         }
      }
   }
   err := json.Unmarshal(d.Token.Raw, &data)
   if err != nil {
      return err
   }
   d.Token.Value = data.Data.Attributes.Token
   return nil
}

type DefaultToken struct {
   Session Value[SessionState]
   Token Value[string]
}

func (d *DefaultToken) New() error {
   req, err := http.NewRequest(
      "", "https://default.any-any.prd.api.discomax.com/token?realm=bolt", nil,
   )
   if err != nil {
      return err
   }
   // fuck you Max
   //req.Header.Set("x-device-info", "!/!(!/!;!/!;!)")
   req.Header.Set("x-device-info", "BEAM-Android/4.12.0 (Google/Android SDK built for x86; ANDROID/7.0; ab3d8214d8d3f368/b6746ddc-7bc7-471f-a16c-f6aaf0c34d26)")
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
   d.Token.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (p *PublicKey) New() error {
   bda, err := get_bda()
   if err != nil {
      return err
   }
   resp, err := http.PostForm(
      "https://wbd-api.arkoselabs.com/fc/gt2/public_key/"+arkose_site_key,
      url.Values{
         "bda": {bda},
         "public_key": {arkose_site_key},
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
   return json.NewDecoder(resp.Body).Decode(p)
}

func get_bda() (string, error) {
   data, err := json.Marshal(map[string][]byte{
      "ct": make([]byte, 3504),
   })
   if err != nil {
      return "", err
   }
   return base64.StdEncoding.EncodeToString(data), nil
}

const (
   arkose_site_key = "B0217B00-2CA4-41CC-925D-1EEB57BFFC2F"
   home_market = "https://default.any-amer.prd.api.discomax.com"
)

func (d *DefaultToken) Playback(web Address) (*Playback, error) {
   body, err := func() ([]byte, error) {
      var p playback_request
      p.ConsumptionType = "streaming"
      p.EditId = web.EditId
      return json.Marshal(p)
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
      "authorization": {"Bearer " + d.Token.Value},
      "content-type": {"application/json"},
      "x-wbd-session-state": {d.Session.Value.String()},
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
   play := &Playback{}
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

func (d *DefaultToken) Routes(web Address) (*DefaultRoutes, error) {
   req, err := http.NewRequest("", home_market, nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      text, _ := web.MarshalText()
      var b strings.Builder
      b.WriteString("/cms/routes")
      b.Write(text)
      return b.String()
   }()
   req.URL.RawQuery = url.Values{
      "include": {"default"},
      // this is not required, but results in a smaller response
      "page[items.size]": {"1"},
   }.Encode()
   req.Header = http.Header{
      "authorization": {"Bearer " + d.Token.Value},
      "x-wbd-session-state": {d.Session.Value.String()},
   }
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
   route := &DefaultRoutes{}
   err = json.NewDecoder(resp.Body).Decode(route)
   if err != nil {
      return nil, err
   }
   return route, nil
}

func (d *DefaultToken) decision() (*default_decision, error) {
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
   req.Header.Set("authorization", "Bearer " + d.Token.Value)
   req.URL.Path = "/labs/api/v1/sessions/feature-flags/decisions"
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   decision := &default_decision{}
   err = json.NewDecoder(resp.Body).Decode(decision)
   if err != nil {
      return nil, err
   }
   return decision, nil
}

func (s SessionState) Set(text string) error {
   for text != "" {
      var key string
      key, text, _ = strings.Cut(text, ";")
      key, value, _ := strings.Cut(key, ":")
      s[key] = value
   }
   return nil
}

func (s SessionState) String() string {
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

type SessionState map[string]string

func (s SessionState) Delete() {
   for key := range s {
      switch key {
      case "device", "token", "user":
      default:
         delete(s, key)
      }
   }
}

type Value[T any] struct {
   Value T
   Raw []byte
}

type DefaultLogin struct {
   Credentials struct {
      Username string `json:"username"`
      Password string `json:"password"`
   } `json:"credentials"`
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

type PublicKey struct {
   Token string
}

type RouteInclude struct {
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

func (p *Playback) RequestUrl() (string, bool) {
   return p.Drm.Schemes.Widevine.LicenseUrl, true
}

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

// note you can use other keys, but you need to change home_market to match
var default_key = hmac_key{
   Id:  "android1_prd",
   Key: []byte("6fd2c4b9-7b43-49ee-a62e-57ffd7bdfe9c"),
}

type hmac_key struct {
   Id  string
   Key []byte
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

func (d DefaultRoutes) video() (*RouteInclude, bool) {
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

type Playback struct {
   Drm struct {
      Schemes struct {
         Widevine struct {
            LicenseUrl string
         }
      }
   }
   Fallback struct {
      Manifest struct {
         Url Manifest
      }
   }
}

type DefaultRoutes struct {
   Data struct {
      Attributes struct {
         Url Address
      }
   }
   Included []RouteInclude
}

func (m *Manifest) UnmarshalText(text []byte) error {
   m.Url = strings.Replace(string(text), "_fallback", "", 1)
   return nil
}

func (a *Address) MarshalText() ([]byte, error) {
   var b bytes.Buffer
   if a.VideoId != "" {
      b.WriteString("/video/watch/")
      b.WriteString(a.VideoId)
   }
   if a.EditId != "" {
      b.WriteByte('/')
      b.WriteString(a.EditId)
   }
   return b.Bytes(), nil
}

type Manifest struct {
   Url string
}

type Address struct {
   EditId  string
   VideoId string
}

func (a *Address) UnmarshalText(text []byte) error {
   s := string(text)
   if !strings.Contains(s, "/video/watch/") {
      return errors.New("/video/watch/ not found")
   }
   s = strings.TrimPrefix(s, "https://")
   s = strings.TrimPrefix(s, "play.max.com")
   s = strings.TrimPrefix(s, "/video/watch/")
   var found bool
   a.VideoId, a.EditId, found = strings.Cut(s, "/")
   if !found {
      return errors.New("/ not found")
   }
   return nil
}
