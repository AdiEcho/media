package roku

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "strings"
)

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

type AccountCode struct {
   Code string
}

func (a *AccountCode) Unmarshal(data []byte) error {
   return json.Unmarshal(data, a)
}

func (a *AccountAuth) Code(data *[]byte) (*AccountCode, error) {
   body, err := json.Marshal(map[string]string{"platform": "googletv"})
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
   body, err = io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   if data != nil {
      *data = body
      return nil, nil
   }
   var code AccountCode
   err = code.Unmarshal(body)
   if err != nil {
      return nil, err
   }
   return &code, nil
}
