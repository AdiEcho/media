package roku

import (
   "bytes"
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strings"
)

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
   play := new(Playback)
   err = json.NewDecoder(resp.Body).Decode(play)
   if err != nil {
      return nil, err
   }
   return play, nil
}

func (a AccountToken) Code() (*ActivationCode, error) {
   body, err := json.Marshal(map[string]string{
      "platform": "googletv",
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://googletv.web.roku.com/api/v1/account/activation",
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
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   code := ActivationCode{Account: a}
   err = json.NewDecoder(resp.Body).Decode(&code.V)
   if err != nil {
      return nil, err
   }
   return &code, nil
}

type ActivationCode struct {
   Account AccountToken
   V     struct {
      Code string
   }
}

func (a ActivationCode) String() string {
   var b strings.Builder
   b.WriteString("1 Visit the URL\n")
   b.WriteString("  therokuchannel.com/link\n")
   b.WriteString("\n")
   b.WriteString("2 Enter the activation code\n")
   b.WriteString("  ")
   b.WriteString(a.V.Code)
   return b.String()
}

func (a *ActivationCode) Unmarshal(text []byte) error {
   return json.Unmarshal(text, a)
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
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   data, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return &ActivationToken{Data: data}, nil
}

func (a *ActivationToken) Unmarshal() error {
   return json.Unmarshal(a.Data, &a.V)
}

type ActivationToken struct {
   Data []byte
   V    struct {
      Token string
   }
}

type Playback struct {
   Drm struct {
      Widevine struct {
         LicenseServer string
      }
   }
   Url string
}

func (Playback) RequestHeader() (http.Header, error) {
   return http.Header{}, nil
}

func (p Playback) RequestUrl() (string, bool) {
   return p.Drm.Widevine.LicenseServer, true
}

func (Playback) WrapRequest(b []byte) ([]byte, error) {
   return b, nil
}

func (Playback) UnwrapResponse(b []byte) ([]byte, error) {
   return b, nil
}
func (a ActivationCode) Marshal() ([]byte, error) {
   return json.MarshalIndent(a, "", " ")
}

const user_agent = "trc-googletv; production; 0"

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
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(a)
}

type AccountToken struct {
   AuthToken string
}
