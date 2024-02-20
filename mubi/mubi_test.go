package mubi

import (
   "154.pages.dev/encoding"
   "fmt"
   "os"
   "testing"
)

// mubi.com/films/325455/player
// mubi.com/films/passages-2022
const passages_2022 = 325455

func TestSecure(t *testing.T) {
   var (
      auth Authenticate
      err error
   )
   auth.Raw, err = os.ReadFile("authenticate.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   var film FilmResponse
   film.s.ID = passages_2022
   secure, err := auth.URL(&film)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", secure)
}

// mubi.com/films/190/player
// mubi.com/films/dogville
var dogvilles = []string{
   "/films/dogville",
   "/en/us/films/dogville",
   "/us/films/dogville",
   "/en/films/dogville",
}

func TestFilm(t *testing.T) {
   for i, dogville := range dogvilles {
      var web WebAddress
      err := web.Set(dogville)
      if err != nil {
         t.Fatal(err)
      }
      if i == 0 {
         film, err := web.Film()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Println(encoding.Name(film))
      }
      fmt.Println(web)
   }
}

func TestAuthenticate(t *testing.T) {
   var (
      code LinkCode
      err error
   )
   code.Raw, err = os.ReadFile("code.json")
   if err != nil {
      t.Fatal(err)
   }
   code.Unmarshal()
   auth, err := code.Authenticate()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.json", auth.Raw, 0666)
}

func TestCode(t *testing.T) {
   var code LinkCode
   err := code.New()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.json", code.Raw, 0666)
   code.Unmarshal()
   fmt.Println(code)
}
