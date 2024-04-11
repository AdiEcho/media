package mubi

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strconv"
   "strings"
)

type TextTrack struct {
   ID string
   URL string
}

func (t TextTrack) String() string {
   var b strings.Builder
   b.WriteString("id = ")
   b.WriteString(t.ID)
   return b.String()
}

type SecureUrl struct {
   Data []byte
   V struct {
      Text_Track_URLs []TextTrack
      URL string
   }
}

func (s *SecureUrl) Unmarshal() error {
   return json.Unmarshal(s.Data, &s.V)
}

func (a Authenticate) URL(f *FilmResponse) (*SecureUrl, error) {
   address := func() string {
      b := []byte("https://api.mubi.com/v3/films/")
      b = strconv.AppendInt(b, f.V.ID, 10)
      b = append(b, "/viewing/secure_url"...)
      return string(b)
   }
   req, err := http.NewRequest("GET", address(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.V.Token},
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
