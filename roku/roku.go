package roku

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
)

func (a *ActivationToken) Unmarshal() error {
   return json.Unmarshal(a.Data, &a.V)
}

type ActivationToken struct {
   Data []byte
   V    struct {
      Token string
   }
}

func (a ActivationCode) Token() (*ActivationToken, error) {
   req, err := http.NewRequest("", "https://googletv.web.roku.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = "/api/v1/account/activation/" + a.V.Code
   req.Header = http.Header{
      "user-agent":           {user_agent},
      "x-roku-content-token": {a.Account.AuthToken},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var token ActivationToken
   token.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &token, nil
}

// token can be nil
func (a *AccountToken) New(token *ActivationToken) error {
   req, err := http.NewRequest("", "https://googletv.web.roku.com", nil)
   if err != nil {
      return err
   }
   req.URL.Path = "/api/v1/account/token"
   req.Header.Set("user-agent", user_agent)
   if token != nil {
      req.Header.Set("x-roku-content-token", token.V.Token)
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
