package kanopy

import (
   "bytes"
   "encoding/json"
   "net/http"
)

type web_token struct {
   Jwt string
}

func (w *web_token) New(email, password string) error {
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
      return err
   }
   req, err := http.NewRequest(
      "POST", "https://www.kanopy.com/kapi/login", bytes.NewReader(data),
   )
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "content-type": {"application/json"},
      "user-agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/115.0"},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer resp.Body.Close()
   return json.NewDecoder(resp.Body).Decode(w)
}
