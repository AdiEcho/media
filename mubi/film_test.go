package mubi

import (
   "154.pages.dev/media/internal"
   "fmt"
   "testing"
)

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
         fmt.Println(internal.Name(Namer{film}))
      }
      fmt.Println(web)
   }
}
