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
   var body struct {
      Query     string `json:"query"`
      Variables struct {
         Email    string `json:"email"`
         Password string `json:"password"`
      } `json:"variables"`
   }
   body.Query = query_user
   body.Variables.Email = email
   body.Variables.Password = password
   raw, err := json.Marshal(body)
   if err != nil {
      return err
   }
   resp, err := http.Post(
      "https://api.audienceplayer.com/graphql/2/user",
      "application/json", bytes.NewReader(raw),
   )
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   o.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   return nil
}

type OperationUser struct {
   AccessToken string `json:"access_token"`
   Raw []byte `json:"-"`
}

func (o *OperationUser) Unmarshal() error {
   var body struct {
      Data struct {
         UserAuthenticate OperationUser
      }
   }
   err := json.Unmarshal(o.Raw, &body)
   if err != nil {
      return err
   }
   *o = body.Data.UserAuthenticate
   return nil
}
