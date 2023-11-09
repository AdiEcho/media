package hulu

import (
   "bytes"
   "encoding/json"
   "net/http"
)

type details struct {
   Items []struct {
      Series_Name string
      Episode_Name string
   }
}

func (a authenticate) details(eab string) (*details, error) {
   body, err := func() ([]byte, error) {
      m := map[string][]string{
         "eabs": {eab},
      }
      return json.Marshal(m)
   }()
   if err != nil {
      return nil, err
   }
   res, err := http.Post(
      "https://guide.hulu.com/guide/details?user_token=" + a.Data.User_Token,
      "application/json",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   detail := new(details)
   if err := json.NewDecoder(res.Body).Decode(detail); err != nil {
      return nil, err
   }
   return detail, nil
}
