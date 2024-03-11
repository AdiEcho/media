package web

import (
   "encoding/json"
   "net/http"
   "strings"
)

type seektable struct {
   PSSH []byte
}

func (s *seektable) New() error {
   address := func() string {
      var b strings.Builder
      b.WriteString("https://seektables.scdn.co")
      b.WriteString("/seektable/392482fe9bed7372d1657d7e22f32b792902f3bd.json")
      return b.String()
   }()
   res, err := http.Get(address)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   return json.NewDecoder(res.Body).Decode(s)
}
