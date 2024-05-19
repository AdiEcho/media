package criterion

import (
   "encoding/json"
   "io"
   "net/http"
   "net/url"
)

const client_id = "9a87f110f79cd25250f6c7f3a6ec8b9851063ca156dae493bf362a7faf146c78"

type AuthToken struct {
   data []byte
   v    struct {
      AccessToken string `json:"access_token"`
   }
}

func (a *AuthToken) New(username, password string) error {
   res, err := http.PostForm("https://auth.vhx.com/v1/oauth/token", url.Values{
      "client_id":  {client_id},
      "grant_type": {"password"},
      "password":   {password},
      "username":   {username},
   })
   if err != nil {
      return err
   }
   defer res.Body.Close()
   a.data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a *AuthToken) unmarshal() error {
   return json.Unmarshal(a.data, &a.v)
}
