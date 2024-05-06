package draken

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
)

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
   a.data, err = io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   return nil
}

func (a *auth_login) unmarshal() error {
   return json.Unmarshal(a.data, &a.v)
}

type auth_login struct {
   data []byte
   v struct {
      Token string
   }
}
