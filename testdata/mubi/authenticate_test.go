package mubi

import (
   "os"
   "testing"
)

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
