package mubi

import (
   "encoding/json"
   "errors"
   "net/http"
   "strconv"
   "strings"
)

func (a Authenticate) URL(f *FilmResponse) (*SecureUrl, error) {
   address := func() string {
      b := []byte("https://api.mubi.com/v3/films/")
      b = strconv.AppendInt(b, f.s.ID, 10)
      b = append(b, "/viewing/secure_url"...)
      return string(b)
   }
   req, err := http.NewRequest("GET", address(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.s.Token},
      "Client": {client},
      "Client-Country": {ClientCountry},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return nil, errors.New(b.String())
   }
   secure := new(SecureUrl)
   if err := json.NewDecoder(res.Body).Decode(secure); err != nil {
      return nil, err
   }
   return secure, nil
}

type SecureUrl struct {
   URL string
}
