package member

import (
   "bytes"
   "encoding/json"
   "net/http"
)

const user_authenticate = `
mutation($email: String, $password: String) {
   UserAuthenticate(email: $email, password: $password) {
      access_token
   }
}
`

type authenticate struct {
   Data struct {
      UserAuthenticate struct {
         AccessToken string `json:"access_token"`
      }
   }
}

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
   return json.NewDecoder(res.Body).Decode(a)
}
