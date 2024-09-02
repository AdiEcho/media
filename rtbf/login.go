package rtbf

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "net/url"
)

// hard coded in JavaScript
const api_key = "4_Ml_fJ47GnBAW6FrPzMxh0w"

type AuvioLogin struct {
   CookieValue string
   Raw []byte
}

func (a *AuvioLogin) New(id, password string) error {
   resp, err := http.PostForm(
      "https://login.auvio.rtbf.be/accounts.login", url.Values{
         "APIKey":   {api_key},
         "loginID":  {id},
         "password": {password},
      },
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   a.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a *AuvioLogin) Token() (*WebToken, error) {
   resp, err := http.PostForm(
      "https://login.auvio.rtbf.be/accounts.getJWT", url.Values{
         "APIKey": {api_key},
         "login_token": {a.CookieValue},
      },
   )
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

func (a *AuvioLogin) Unmarshal() error {
   var data struct {
      ErrorMessage string
      SessionInfo  struct {
         CookieValue string
      }
   }
   err := json.Unmarshal(a.Raw, &data)
   if err != nil {
      return err
   }
   if v := data.ErrorMessage; v != "" {
      return errors.New(v)
   }
   a.CookieValue = data.SessionInfo.CookieValue
   return nil
}
