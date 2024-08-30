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
   auth.Data, err = os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   secure, err := auth.Url(&FilmResponse{Id: test.id})
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", secure)
}
