package rtbf

import (
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "strings"
)

// hard coded in JavaScript
const api_key = "4_Ml_fJ47GnBAW6FrPzMxh0w"

func (o one) login(id, password string) (*accounts_login, error) {
   body := url.Values{
      "APIKey": {api_key},
      "loginID": {id},
      "password": {password},
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
   var login accounts_login
   login.data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &login, nil
}

func (a *accounts_login) unmarshal() error {
   return json.Unmarshal(a.data, &a.v)
}

type accounts_login struct {
   data []byte
   v struct {
      SessionInfo struct {
         CookieValue string
      }
   }
}
