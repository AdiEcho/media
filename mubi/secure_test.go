package mubi

import (
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
   auth.Data, err = os.ReadFile("authenticate.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   var film FilmResponse
   film.V.ID = passages_2022
   secure, err := auth.URL(&film)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", secure)
}
