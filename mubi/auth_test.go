package mubi

import (
   "fmt"
   "os"
   "reflect"
   "testing"
)

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Address{},
   Authenticate{},
   FilmResponse{},
   LinkCode{},
   Namer{},
   SecureUrl{},
   TextTrack{},
}

func TestAuthenticate(t *testing.T) {
   var (
      code LinkCode
      err error
   )
   code.Raw, err = os.ReadFile("code.txt")
   if err != nil {
      t.Fatal(err)
   }
   code.Unmarshal()
   auth, err := code.Authenticate()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.txt", auth.Raw, os.ModePerm)
}
