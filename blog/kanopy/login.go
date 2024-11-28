package kanopy

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
)

const user_agent = "!"

func (web_token) marshal(email, password string) ([]byte, error) {
   var value struct {
      CredentialType string `json:"credentialType"`
      User struct {
         Email    string `json:"email"`
         Password string `json:"password"`
      } `json:"emailUser"`
   }
   value.CredentialType = "email"
   value.User.Email = email
   value.User.Password = password
   data, err := json.Marshal(value)
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://www.kanopy.com/kapi/login", bytes.NewReader(data),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type": {"application/json"},
      "user-agent": {user_agent},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   return io.ReadAll(resp.Body)
}

// good for 10 years
type web_token struct {
   Jwt string
}

func (w *web_token) unmarshal(data []byte) error {
   return json.Unmarshal(data, w)
}
