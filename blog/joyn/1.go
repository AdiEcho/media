package joyn

import (
   "bytes"
   "encoding/json"
   "net/http"
)

func anonymous() (*http.Response, error) {
   body, err := func() ([]byte, error) {
      m := map[string]string{
         "client_id": "!",
         "client_name": "web",
      }
      return json.Marshal(m)
   }()
   if err != nil {
      return nil, err
   }
   return http.Post(
      "https://auth.joyn.de/auth/anonymous", "application/json",
      bytes.NewReader(body),
   )
}
