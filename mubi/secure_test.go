package mubi

import (
   "fmt"
   "os"
   "testing"
)

func TestSecure(t *testing.T) {
   data, err := os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   var auth Authenticate
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   var secure SecureUrl
   data, err = secure.Marshal(&auth, &FilmResponse{Id: test.id})
   if err != nil {
      t.Fatal(err)
   }
   err = secure.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", secure)
}
