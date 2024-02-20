package mubi

import (
   "encoding/json"
   "net/http"
   "strconv"
)

type secure_url struct {
   URL string
}

func (a authenticate) secure(film int64) (*secure_url, error) {
   address := func() string {
      b := []byte("https://api.mubi.com/v3/films/")
      b = strconv.AppendInt(b, film, 10)
      b = append(b, "/viewing/secure_url"...)
      return string(b)
   }
   req, err := http.NewRequest("GET", address(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.s.Token},
      "Client": {"web"},
      "Client-Country": {client_country},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   secure := new(secure_url)
   if err := json.NewDecoder(res.Body).Decode(secure); err != nil {
      return nil, err
   }
   return secure, nil
}
