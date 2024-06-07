package rtbf

import (
   "encoding/json"
   "errors"
   "net/http"
   "net/url"
   "strings"
)

func (a accounts_login) token() (*web_token, error) {
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
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   var web web_token
   err = json.NewDecoder(res.Body).Decode(&web)
   if err != nil {
      return nil, err
   }
   if v := web.ErrorMessage; v != "" {
      return nil, errors.New(v)
   }
   return &web, nil
}

type web_token struct {
   ErrorMessage string
   IdToken string `json:"id_token"`
}
