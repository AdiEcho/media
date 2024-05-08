package joyn

import (
   "bytes"
   "encoding/json"
   "net/http"
)

type anonymous struct {
   Access_Token string
}

func (a *anonymous) New() error {
   body, err := func() ([]byte, error) {
      m := map[string]string{
         "client_id": "!",
         "client_name": "web",
      }
      return json.Marshal(m)
   }()
   if err != nil {
      return err
   }
   res, err := http.Post(
      "https://auth.joyn.de/auth/anonymous", "application/json",
      bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(a)
}
