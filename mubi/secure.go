package mubi

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strconv"
   "strings"
)

type SecureUrl struct {
   Data []byte
   V struct {
      URL string
   }
}

func (s *SecureUrl) Unmarshal() error {
   return json.Unmarshal(s.Data, &s.V)
}

func (a Authenticate) URL(f *FilmResponse) (*SecureUrl, error) {
   address := func() string {
      b := []byte("https://api.mubi.com/v3/films/")
      b = strconv.AppendInt(b, f.v.ID, 10)
      b = append(b, "/viewing/secure_url"...)
      return string(b)
   }
   req, err := http.NewRequest("GET", address(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.v.Token},
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
   var secure SecureUrl
   secure.Data, err = io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   return &secure, nil
}
