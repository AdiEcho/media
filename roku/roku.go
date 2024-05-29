package roku

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

type activation_token struct {
   data []byte
   v struct {
      Token string
   }
}

func (a activation_code) token() (*activation_token, error) {
   req, err := http.NewRequest("", "https://googletv.web.roku.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/v1/account/activation/" + a.V.Code
   req.Header = http.Header{
      "user-agent": {user_agent},
      "x-roku-content-token": {a.Token.AuthToken},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var token activation_token
   token.data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &token, nil
}

func (a *activation_token) unmarshal() error {
   return json.Unmarshal(a.data, &a.v)
}
const user_agent = "Mozilla/5.0 (Linux; Android 9; sdk_google_atv_x86 Build/PSR1.180720.121; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.158 Mobile Safari/537.36 googletv; trc-googletv; production; 0.f901664681ba61e2"

type account_token struct {
   AuthToken string
}

// token can be nil
func (a *account_token) New(token *activation_token) error {
   req, err := http.NewRequest("", "https://googletv.web.roku.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/api/v1/account/token"
   req.Header.Set("user-agent", user_agent)
   if token != nil {
      req.Header.Set("x-roku-content-token", token.v.Token)
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
type AccountToken struct {
   AuthToken string
}

func (a *AccountToken) New() error {
   req, err := http.NewRequest(
      "GET", "https://googletv.web.roku.com/api/v1/account/token", nil,
   )
   if err != nil {
      return err
   }
   req.Header.Set("user-agent", user_agent)
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}

func (a AccountToken) Playback(roku_id string) (*Playback, error) {
   body, err := json.Marshal(map[string]string{
      "mediaFormat": "DASH",
      "providerId":  "rokuavod",
      "rokuId":      roku_id,
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://googletv.web.roku.com/api/v3/playback",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type":         {"application/json"},
      "user-agent":           {user_agent},
      "x-roku-content-token": {a.AuthToken},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errors.New(res.Status)
   }
   play := new(Playback)
   err = json.NewDecoder(res.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

type Playback struct {
   DRM struct {
      Widevine struct {
         LicenseServer string
      }
   }
   URL string
}

const user_agent = "trc-googletv; production; 0"

func (Playback) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (p Playback) RequestUrl() (string, bool) {
   return p.DRM.Widevine.LicenseServer, true
}

func (Playback) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}
