package mubi

import (
   "fmt"
   "os"
   "testing"
)

func TestSecure(t *testing.T) {
   var (
      auth Authenticate
      err error
   )
   auth.Raw, err = os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = auth.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   secure, err := auth.Url(&FilmResponse{Id: test.id})
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", secure)
}
