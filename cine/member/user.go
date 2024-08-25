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

type OperationUser struct {
   AccessToken string `json:"access_token"`
   Raw []byte `json:"-"`
}

func (o *OperationUser) New(email, password string) error {
   var value struct {
      Query     string `json:"query"`
      Variables struct {
         Email    string `json:"email"`
         Password string `json:"password"`
      } `json:"variables"`
   }
   value.Query = query_user
   value.Variables.Email = email
   value.Variables.Password = password
   data, err := json.Marshal(value)
   if err != nil {
      return err
   }
   resp, err := http.Post(
      "https://api.audienceplayer.com/graphql/2/user",
      "application/json", bytes.NewReader(data),
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

func (o *OperationUser) Unmarshal() error {
   var value struct {
      Data struct {
         UserAuthenticate OperationUser
      }
   }
   err := json.Unmarshal(o.Raw, &value)
   if err != nil {
      return err
   }
   *o = value.Data.UserAuthenticate
   return nil
}
