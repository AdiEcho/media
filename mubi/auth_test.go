package mubi

import (
   "fmt"
   "os"
   "reflect"
   "testing"
)

func TestAuthenticate(t *testing.T) {
   data, err := os.ReadFile("code.txt")
   if err != nil {
      t.Fatal(err)
   }
   var code LinkCode
   err = code.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   data, err = (*Authenticate).Marshal(nil, &code)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.txt", data, os.ModePerm)
}

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
