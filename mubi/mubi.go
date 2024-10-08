package mubi

import (
   "encoding/json"
   "errors"
   "net/http"
   "strings"
)

var ClientCountry = "US"

// "android" requires headers:
// client-device-identifier
// client-version
const client = "web"

type Address struct {
   s string
}

func (a *Address) Set(text string) error {
   var ok bool
   _, a.s, ok = strings.Cut(text, "/films/")
   if !ok {
      return errors.New("/films/")
   }
   return nil
}

func (Namer) Episode() int {
   return 0
}

func (Namer) Season() int {
   return 0
}

func (Namer) Show() string {
   return ""
}

type TextTrack struct {
   Id string
   Url string
}

type FilmResponse struct {
   Id int64
   Title string
   Year int
}

func (a *Address) String() string {
   return a.s
}

func (n *Namer) Title() string {
   return n.Film.Title
}

func (n *Namer) Year() int {
   return n.Film.Year
}

type Namer struct {
   Film *FilmResponse
}

func (t *TextTrack) String() string {
   return "id = " + t.Id
}

func (a *Address) Film() (*FilmResponse, error) {
   req, err := http.NewRequest("", "https://api.mubi.com/v3/films/" + a.s, nil)
   if err != nil {
      return nil, err
   }
   req.Header = http.Header{
      "client": {client},
      "client-country": {ClientCountry},
   }
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      return nil, err
   }
   defer resp.Body.Close()
   film := &FilmResponse{}
   err = json.NewDecoder(resp.Body).Decode(film)
   if err != nil {
      return nil, err
   }
   return film, nil
}
