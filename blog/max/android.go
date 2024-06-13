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

func (default_token) playback(p playback_request) (*http.Response, error) {
   body, err := json.Marshal(p)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://default.any-any.prd.api.max.com", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/any/playback/v1/playbackInfo"
   req.Header.Set("content-type", "application/json")
   // req.AddCookie(st.Cookie)
   return http.DefaultClient.Do(req)
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
   req.URL.Path = "/labs/api/v1/sessions/feature-flags/decisions"
   req.Header.Set("authorization", "Bearer " + d.Data.Attributes.Token)
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

var android_config = key_config{
   Id: "android1_prd",
   Key: []byte("6fd2c4b9-7b43-49ee-a62e-57ffd7bdfe9c"),
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
