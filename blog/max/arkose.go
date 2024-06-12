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

func (st st_cookie) marshal() ([]byte, error) {
   return json.MarshalIndent(st, "", " ")
}

func (st *st_cookie) unmarshal(text []byte) error {
   return json.Unmarshal(text, st)
}

type st_cookie struct {
   Cookie *http.Cookie
}

type key_config struct {
   Id string
   Key []byte
}

var config = key_config{
   Id: "web1_prd",
   Key: []byte("9d697c21-2ec9-494b-a90d-e3de471e6e9f"),
}

type public_key struct {
   Token string
}

const arkose_site_key = "B0217B00-2CA4-41CC-925D-1EEB57BFFC2F"

type default_login struct {
   Credentials struct {
      Username string `json:"username"`
      Password string `json:"password"`
   } `json:"credentials"`
}

////////////

func (st *st_cookie) login(key public_key, login default_login) error {
   body, err := json.Marshal(login)
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://default.any-amer.prd.api.max.com/login",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.AddCookie(st.Cookie)
   req.Header.Set("content-type", "application/json")
   req.Header.Set("x-disco-arkose-token", key.Token)
   req.Header.Set("x-disco-client-id", func() string {
      timestamp := time.Now().Unix()
      hash := hmac.New(sha256.New, config.Key)
      fmt.Fprintf(hash, "%v:POST:/login:%s", timestamp, body)
      signature := hash.Sum(nil)
      return fmt.Sprintf("%v:%v:%x", config.Id, timestamp, signature)
   }())
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   for _, cookie := range resp.Cookies() {
      if cookie.Name == "st" {
         st.Cookie = cookie
         return nil
      }
   }
   return http.ErrNoCookie
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

func (st *st_cookie) New() error {
   req, err := http.NewRequest(
      "", "https://default.any-any.prd.api.max.com/token?realm=bolt", nil,
   )
   if err != nil {
      return err
   }
   req.Header.Set(
      "x-device-info", "beam/4.0.1 (desktop/desktop; Windows/10; !/!)",
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
   for _, cookie := range resp.Cookies() {
      if cookie.Name == "st" {
         st.Cookie = cookie
         return nil
      }
   }
   return http.ErrNoCookie
}

func (st st_cookie) config() (*key_config, error) {
   body, err := json.Marshal(map[string]string{
      "projectId": "67e7aa0f-b186-4b85-9cb0-86d40a23636c",
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://default.beam-any.prd.api.max.com", bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/labs/api/v1/sessions/feature-flags/decisions"
   req.AddCookie(st.Cookie)
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var decision struct {
      HmacKeys struct {
         Config struct {
            Web key_config
         }
      }
   }
   err = json.NewDecoder(resp.Body).Decode(&decision)
   if err != nil {
      return nil, err
   }
   return &decision.HmacKeys.Config.Web, nil
}
