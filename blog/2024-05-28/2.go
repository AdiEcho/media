package roku

import (
   "bytes"
   "encoding/json"
   "net/http"
   "strings"
)

func (t *two_response) unmarshal(text []byte) error {
   return json.Unmarshal(text, t)
}

func (t two_response) marshal() ([]byte, error) {
   return json.Marshal(t)
}

type two_response struct {
   One one_response
   Two struct {
      Code string
   }
}

func (o one_response) two() (*two_response, error) {
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
      "x-roku-content-token": {o.AuthToken},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   two := two_response{One: o}
   err = json.NewDecoder(res.Body).Decode(&two.Two)
   if err != nil {
      return nil, err
   }
   return &two, nil
}

func (t two_response) String() string {
   var b strings.Builder
   b.WriteString("1 Visit the URL\n")
   b.WriteString("  therokuchannel.com/link\n")
   b.WriteString("\n")
   b.WriteString("2 Enter the activation code\n")
   b.WriteString("  ")
   b.WriteString(t.Two.Code)
   return b.String()
}
