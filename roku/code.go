package roku

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strings"
)

type AccountCode struct {
   Code string
   Raw []byte `json:"-"`
}

func (a *AccountAuth) Code() (*AccountCode, error) {
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
      "content-type":         {"application/json"},
      "user-agent":           {user_agent},
      "x-roku-content-token": {a.AuthToken},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   var code AccountCode
   code.Raw, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   return &code, nil
}

func (a *AccountCode) String() string {
   var b strings.Builder
   b.WriteString("1 Visit the URL\n")
   b.WriteString("  therokuchannel.com/link\n")
   b.WriteString("\n")
   b.WriteString("2 Enter the activation code\n")
   b.WriteString("  ")
   b.WriteString(a.Code)
   return b.String()
}

func (a *AccountCode) Unmarshal() error {
   return json.Unmarshal(a.Raw, a)
}
