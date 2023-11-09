package hulu

import (
   "bytes"
   "encoding/json"
   "net/http"
)

func (a authenticate) details(eab string) (*http.Response, error) {
   body, err := func() ([]byte, error) {
      v := struct{
         Eabs []string
      }{
         []string{eab},
      }
      return json.Marshal(v)
   }()
   if err != nil {
      return nil, err
   }
   return http.Post(
      "https://guide.hulu.com/guide/details?user_token=" + a.Data.User_Token,
      "application/json",
      bytes.NewReader(body),
   )
}
