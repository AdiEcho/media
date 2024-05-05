package draken

import (
   "bytes"
   "encoding/json"
   "net/http"
)

type auth_login struct {
   Token string
}

func (a *auth_login) New(identity, key string) error {
   body, err := func() ([]byte, error) {
      m := map[string]string{
         "identity": identity,
         "accessKey": key,
      }
      return json.Marshal(m)
   }()
   if err != nil {
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://drakenfilm.se/api/auth/login", bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header.Set("content-type", "application/json")
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
