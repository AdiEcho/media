package roku

import (
   "bytes"
   "encoding/json"
   "net/http"
   "strings"
)

type activation_code struct {
   Token account_token
   V struct {
      Code string
   }
}

func (a account_token) code() (*activation_code, error) {
   body, err := json.Marshal(map[string]string{
      "platform": "googletv",
   })
   if err != nil {
      return nil, err
   }
   req, err := http.NewRequest(
      "POST", "https://googletv.web.roku.com/api/v1/account/activation",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "content-type": {"application/json"},
      "user-agent": {user_agent},
      "x-roku-content-token": {a.AuthToken},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   code := activation_code{Token: a}
   err = json.NewDecoder(res.Body).Decode(&code.V)
   if err != nil {
      return nil, err
   }
   return &code, nil
}

func (a activation_code) String() string {
   var b strings.Builder
   b.WriteString("1 Visit the URL\n")
   b.WriteString("  therokuchannel.com/link\n")
   b.WriteString("\n")
   b.WriteString("2 Enter the activation code\n")
   b.WriteString("  ")
   b.WriteString(a.V.Code)
   return b.String()
}

func (a *activation_code) unmarshal(text []byte) error {
   return json.Unmarshal(text, a)
}

func (a activation_code) marshal() ([]byte, error) {
   return json.Marshal(a)
}
