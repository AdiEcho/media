package mubi

import (
   "encoding/json"
   "errors"
   "io"
   "net/http"
   "strconv"
   "strings"
)

func (a Authenticate) URL(film *FilmResponse) (*SecureUrl, error) {
   address := func() string {
      b := []byte("https://api.mubi.com/v3/films/")
      b = strconv.AppendInt(b, film.ID, 10)
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

func (Name) Episode() int {
   return 0
}

func (Name) Season() int {
   return 0
}

func (Name) Show() string {
   return ""
}

func (n Name) Title() string {
   return n.F.Title
}

func (n Name) Year() int {
   return n.F.Year
}

// Mubi do this sneaky thing. you cannot download a video unless you have told
// the API that you are watching it. so you have to call
// `/v3/films/%v/viewing`, otherwise it wont let you get the MPD. if you have
// already viewed the video on the website that counts, but if you only use the
// tool it will error
func (a Authenticate) Viewing(film *FilmResponse) error {
   address := func() string {
      b := []byte("https://api.mubi.com/v3/films/")
      b = strconv.AppendInt(b, film.ID, 10)
      b = append(b, "/viewing"...)
      return string(b)
   }
   req, err := http.NewRequest("POST", address(), nil)
   if err != nil {
      return err
   }
   req.Header = http.Header{
      "Authorization": {"Bearer " + a.V.Token},
      "Client": {client},
      "Client-Country": {ClientCountry},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      var b strings.Builder
      res.Write(&b)
      return errors.New(b.String())
   }
   return nil
}

type FilmResponse struct {
   ID int64
   Title string
   Year int
}

type Name struct {
   F *FilmResponse
}

func (w WebAddress) Film() (*FilmResponse, error) {
   req, err := http.NewRequest(
      "GET", "https://api.mubi.com/v3/films/" + w.s, nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "Client": {client},
      "Client-Country": {ClientCountry},
   }
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   film := new(FilmResponse)
   if err := json.NewDecoder(res.Body).Decode(film); err != nil {
      return nil, err
   }
   return film, nil
}
