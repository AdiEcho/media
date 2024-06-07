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

func (a *accounts_login) New(id, password string) error {
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
      return err
   }
   req.Header.Set("content-type", "application/x-www-form-urlencoded")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   err = json.NewDecoder(res.Body).Decode(a)
   if err != nil {
      return err
   }
   if a.ErrorDetails != "" {
      return errors.New(a.ErrorDetails)
   }
   return nil
}
