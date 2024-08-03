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

func (d *DefaultToken) Marshal() ([]byte, error) {
   return json.Marshal(d)
}

type DefaultToken struct {
   SessionState SessionState
   Body struct {
      Data struct {
         Attributes struct {
            Token string
         }
      }
   }
}

func (d *DefaultToken) Unmarshal(text []byte) error {
   return json.Unmarshal(text, d)
}

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
   d.SessionState = make(SessionState)
   d.SessionState.Set(resp.Header.Get("x-wbd-session-state"))
   return json.NewDecoder(resp.Body).Decode(&d.Body)
}

func (d *DefaultToken) Playback(flag AddressFlag) (*Playback, error) {
   body, err := func() ([]byte, error) {
      var p playback_request
      p.ConsumptionType = "streaming"
      p.EditId = flag.EditId
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
      "authorization": {"Bearer "+d.Body.Data.Attributes.Token},
      "content-type": {"application/json"},
      "x-wbd-session-state": {d.SessionState.String()},
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

func (d *DefaultToken) Routes(flag AddressFlag) (*DefaultRoutes, error) {
   address := func() string {
      path, _ := flag.MarshalText()
      var b strings.Builder
      b.WriteString("https://default.any-")
      b.WriteString(home_market)
      b.WriteString(".prd.api.discomax.com/cms/routes")
      b.Write(path)
      return b.String()
   }()
   req, err := http.NewRequest("", address, nil)
   if err != nil {
      return nil, err
   }
   req.URL.RawQuery = url.Values{
      "include": {"default"},
      // this is not required, but results in a smaller response
      "page[items.size]": {"1"},
   }.Encode()
   req.Header = http.Header{
      "authorization": {"Bearer "+d.Body.Data.Attributes.Token},
      "x-wbd-session-state": {d.SessionState.String()},
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
   route := new(DefaultRoutes)
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

func (s SessionState) Delete() {
   for key := range s {
      switch key {
      case "device", "token", "user":
      default:
         delete(s, key)
      }
   }
}

type SessionState map[string]string
