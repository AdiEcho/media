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

func (r login_request) login(
   key public_key,
   token default_token,
) (*login_response, error) {
   body, err := json.Marshal(r)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://default.any-amer.prd.api.max.com/login",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(token.cookie)
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
      return nil, err
   }
   defer resp.Body.Close()
   login := new(login_response)
   err = json.NewDecoder(resp.Body).Decode(login)
   if err != nil {
      return nil, err
   }
   return login, nil
}

const arkose_site_key = "B0217B00-2CA4-41CC-925D-1EEB57BFFC2F"

func (d default_token) config() (*key_config, error) {
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
   req.AddCookie(d.cookie)
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

type default_token struct {
   cookie *http.Cookie
}

func (d *default_token) New() error {
   req, err := http.NewRequest(
      "", "https://default.any-any.prd.api.max.com/token?realm=bolt", nil,
   )
   if err != nil {
      return err
   }
   req.Header.Set(
      "x-device-info", "beam/4.0.1 (desktop/desktop; Windows/10; !/!)",
   )
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b bytes.Buffer
      res.Write(&b)
      return errors.New(b.String())
   }
   for _, cookie := range res.Cookies() {
      if cookie.Name == "st" {
         d.cookie = cookie
         return nil
      }
   }
   return http.ErrNoCookie
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

type login_request struct {
   Credentials struct {
      Username string `json:"username"`
      Password string `json:"password"`
   } `json:"credentials"`
}

type login_response struct {
   Data struct {
      Attributes struct {
         Token string
      }
   }
}
