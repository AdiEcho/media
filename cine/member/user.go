package member

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
)

const query_user = `
mutation($email: String, $password: String) {
   UserAuthenticate(email: $email, password: $password) {
      access_token
   }
}
`

func (o *OperationUser) New(email, password string) error {
   body, err := func() ([]byte, error) {
      var s struct {
         Query     string `json:"query"`
         Variables struct {
            Email    string `json:"email"`
            Password string `json:"password"`
         } `json:"variables"`
      }
      s.Query = query_user
      s.Variables.Email = email
      s.Variables.Password = password
      return json.Marshal(s)
   }()
   if err != nil {
      return err
   }
   resp, err := http.Post(
      "https://api.audienceplayer.com/graphql/2/user",
      "application/json", bytes.NewReader(body),
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   o.raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

func (o OperationUser) Marshal() []byte {
   return o.raw
}

type OperationUser struct {
   Data *struct {
      UserAuthenticate struct {
         AccessToken string `json:"access_token"`
      }
   }
   raw []byte
}

func (o *OperationUser) Unmarshal(raw []byte) error {
   return json.Unmarshal(raw, o)
}
