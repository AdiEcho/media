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

func (d default_token) decision() (*default_decision, error) {
   body, err := json.Marshal(map[string]string{
      // android1_prd
      "projectId": "d8665e86-8706-415d-8d84-d55ceddccfb5",
      // web1_prd
      //"projectId":"67e7aa0f-b186-4b85-9cb0-86d40a23636c",
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
   decision := new(default_decision)
   err = json.NewDecoder(resp.Body).Decode(decision)
   if err != nil {
      return nil, err
   }
   return decision, nil
}

type default_login struct {
   Credentials struct {
      Username string `json:"username"`
      Password string `json:"password"`
   } `json:"credentials"`
}

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

func (d *default_token) New() error {
   req, err := http.NewRequest(
      "", "https://default.any-any.prd.api.discomax.com/token?realm=bolt", nil,
   )
   if err != nil {
      return err
   }
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
   return json.NewDecoder(resp.Body).Decode(d)
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
   return json.NewDecoder(resp.Body).Decode(d)
}

type hmac_keys struct {
   Android *hmac_key
   AndroidTv *hmac_key
   FireTv *hmac_key
   Hwa *hmac_key
   Ios *hmac_key
   TvOs *hmac_key
   Web *hmac_key
}

type default_decision struct {
   Config struct {
      Config struct {
         HmacKeys hmac_keys
      }
   }
   HmacKeys struct {
      Config hmac_keys
   }
}

type hmac_key struct {
   Id string
   Key []byte
}

var default_key = hmac_key{
   Id: "android1_prd",
   Key: []byte("6fd2c4b9-7b43-49ee-a62e-57ffd7bdfe9c"),
}
