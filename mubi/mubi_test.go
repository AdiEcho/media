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

func TestAuthenticate(t *testing.T) {
   var (
      code linkCode
      err error
   )
   code.Raw, err = os.ReadFile("code.json")
   if err != nil {
      t.Fatal(err)
   }
   code.unmarshal()
   auth, err := code.authenticate()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.json", auth.Raw, 0666)
}

func TestCode(t *testing.T) {
   var code linkCode
   err := code.New()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.json", code.Raw, 0666)
   code.unmarshal()
   fmt.Println(code)
}
