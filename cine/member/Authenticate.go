package member

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
)

const user_authenticate = `
mutation($email: String, $password: String) {
   UserAuthenticate(email: $email, password: $password) {
      access_token
   }
}
`

func (a *authenticate) New(email, password string) error {
   body, err := func() ([]byte, error) {
      var s struct {
         Query string `json:"query"`
         Variables struct {
            Email string `json:"email"`
            Password string `json:"password"`
         } `json:"variables"`
      }
      s.Query = user_authenticate
      s.Variables.Email = email
      s.Variables.Password = password
      return json.Marshal(s)
   }()
   if err != nil {
      return err
   }
   res, err := http.Post(
      "https://api.audienceplayer.com/graphql/2/user",
      "application/json", bytes.NewReader(body),
   )
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

type authenticate struct {
   data []byte
   v struct {
      Data struct {
         UserAuthenticate struct {
            AccessToken string `json:"access_token"`
         }
      }
   }
}

func (a *authenticate) unmarshal() error {
   return json.Unmarshal(a.data, &a.v)
}
