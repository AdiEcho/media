package rtbf

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

// hard coded in JavaScript
const api_key = "4_Ml_fJ47GnBAW6FrPzMxh0w"

func (a AccountLogin) Token() (*WebToken, error) {
   body := url.Values{
      "APIKey": {api_key},
      // from /accounts.login
      "login_token": {a.SessionInfo.CookieValue},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://login.auvio.rtbf.be/accounts.getJWT",
      strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("content-type", "application/x-www-form-urlencoded")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var web WebToken
   err = json.NewDecoder(resp.Body).Decode(&web)
   if err != nil {
      return nil, err
   }
   if v := web.ErrorMessage; v != "" {
      return nil, errors.New(v)
   }
   return &web, nil
}

func (a *AccountLogin) New(id, password string) error {
   body := url.Values{
      "APIKey":   {api_key},
      "loginID":  {id},
      "password": {password},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://login.auvio.rtbf.be/accounts.login",
      strings.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header.Set("content-type", "application/x-www-form-urlencoded")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   err = json.NewDecoder(resp.Body).Decode(a)
   if err != nil {
      return err
   }
   if v := a.ErrorMessage; v != "" {
      return errors.New(v)
   }
   return nil
}

type AccountLogin struct {
   ErrorMessage string
   SessionInfo  struct {
      CookieValue string
   }
}

func (a *AccountLogin) Unmarshal(text []byte) error {
   return json.Unmarshal(text, a)
}

func (a AccountLogin) Marshal() ([]byte, error) {
   return json.Marshal(a)
}
