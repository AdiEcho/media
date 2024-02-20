package mubi

import (
   "net/http"
   "strconv"
)

func (a authenticate) secure_url(film int64) (*http.Response, error) {
   address := func() string {
      var b []byte
      b = append(b, "https://api.mubi.com/v3/films/"...)
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
   return http.DefaultClient.Do(req)
}
