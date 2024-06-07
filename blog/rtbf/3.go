package rtbf

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (a *accounts_login) unmarshal(text []byte) error {
   return json.Unmarshal(text, a)
}

func (a accounts_login) marshal() ([]byte, error) {
   return json.Marshal(a)
}

type accounts_login struct {
   ErrorDetails string
   SessionInfo struct {
      CookieValue string
   }
}

// hard coded in JavaScript
const api_key = "4_Ml_fJ47GnBAW6FrPzMxh0w"

func (o one) login(id, password string) (*accounts_login, error) {
   body := url.Values{
      "loginID": {id},
      "password": {password},
      "APIKey": {api_key},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://login.auvio.rtbf.be/accounts.login",
      strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.AddCookie(o.cookie)
   req.Header.Set("content-type", "application/x-www-form-urlencoded")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   login := new(accounts_login)
   err = json.NewDecoder(res.Body).Decode(login)
   if err != nil {
      return nil, err
   }
   if login.ErrorDetails != "" {
      return nil, errors.New(login.ErrorDetails)
   }
   return login, nil
}
