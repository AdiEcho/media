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
   Data struct {
      UserAuthenticate struct {
         AccessToken string `json:"access_token"`
      }
   }
}

func (o *OperationUser) Unmarshal(data []byte) error {
   return json.Unmarshal(data, o)
}

func (o *OperationUser) New(email, password string, data *[]byte) error {
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
   body, err := json.Marshal(value)
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
   body, err = io.ReadAll(resp.Body)
   if err != nil {
      return err
   }
   if data != nil {
      *data = body
      return nil
   }
   return o.Unmarshal(body)
}
