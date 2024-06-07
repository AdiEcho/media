package rtbf

import (
   "encoding/json"
   "net/http"
   "net/url"
   "strings"
)

type web_token struct {
   IdToken string `json:"id_token"`
}

func (o one) four(login *accounts_login) (*web_token, error) {
   body := url.Values{
      "APIKey": {api_key},
      // from /accounts.login
      "login_token": {login.SessionInfo.CookieValue},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://login.auvio.rtbf.be/accounts.getJWT",
      strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("content-type", "application/x-www-form-urlencoded")
   req.AddCookie(o.cookie) // from /accounts.webSdkBootstrap
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   web := new(web_token)
   err = json.NewDecoder(res.Body).Decode(web)
   if err != nil {
      return nil, err
   }
   return web, nil
}
