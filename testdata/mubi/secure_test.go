package mubi

import (
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
   res, err := auth.secure_url(passages_2022)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
