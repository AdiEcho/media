package web

import (
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

type seektable struct {
   PSSH []byte
}

func (s *seektable) New(file_id string) error {
   address := func() string {
      var b strings.Builder
      b.WriteString("https://seektables.scdn.co")
      b.WriteString("/seektable/")
      b.WriteString(file_id)
      b.WriteString(".json")
      return b.String()
   }()
   res, err := http.Get(address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return errors.New(b.String())
   }
   return json.NewDecoder(res.Body).Decode(s)
}
