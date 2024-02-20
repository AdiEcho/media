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
      auth authenticate
      err error
   )
   auth.Raw, err = os.ReadFile("authenticate.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.unmarshal()
   secure, err := auth.secure(passages_2022)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", secure)
}
