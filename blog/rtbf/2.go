package rtbf

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

// hard coded in JavaScript
const api_key = "4_Ml_fJ47GnBAW6FrPzMxh0w"

type accounts_login struct {
   SessionInfo struct {
      CookieValue string
   }
}

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
   login := new(accounts_login)
   err = json.NewDecoder(res.Body).Decode(login)
   if err != nil {
      return nil, err
   }
   return login, nil
}
