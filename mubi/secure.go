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
   TextTrackUrls []TextTrack `json:"text_track_urls"`
   Url string
}

func (s *SecureUrl) Unmarshal(data []byte) error {
   return json.Unmarshal(data, s)
}

func (a *Authenticate) Secure(
   film *FilmResponse, data *[]byte,
) (*SecureUrl, error) {
   req, err := http.NewRequest("", "https://api.mubi.com", nil)
   if err != nil {
      return nil, err
   }
   req.URL.Path = func() string {
      b := []byte("/v3/films/")
      b = strconv.AppendInt(b, film.Id, 10)
      b = append(b, "/viewing/secure_url"...)
      return string(b)
   }()
   req.Header = http.Header{
      "authorization": {"Bearer " + a.Token},
      "client": {client},
      "client-country": {ClientCountry},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      var b strings.Builder
      resp.Write(&b)
      return nil, errors.New(b.String())
   }
   body, err := io.ReadAll(resp.Body)
   if err != nil {
      return nil, err
   }
   if data != nil {
      *data = body
      return nil, nil
   }
   var secure SecureUrl
   err = secure.Unmarshal(body)
   if err != nil {
      return nil, err
   }
   return &secure, nil
}
