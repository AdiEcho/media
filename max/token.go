package max

import (
   "bytes"
   "crypto/hmac"
   "crypto/sha256"
   "encoding/json"
   "errors"
   "fmt"
   "net/http"
   "time"
)

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
   d.SessionState = resp.Header.Get("x-wbd-session-state")
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

// note you can use other keys, but you need to change home_market to match
var default_key = hmac_key{
   Id:  "android1_prd",
   Key: []byte("6fd2c4b9-7b43-49ee-a62e-57ffd7bdfe9c"),
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
