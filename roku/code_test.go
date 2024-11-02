package roku

import (
   "fmt"
   "os"
   "reflect"
   "testing"
)

func TestCode(t *testing.T) {
   // AccountAuth
   var auth AccountAuth
   data, err := auth.Marshal(nil)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("auth.txt", data, os.ModePerm)
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   // AccountCode
   var code AccountCode
   data, err = code.Marshal(&auth)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.txt", data, os.ModePerm)
   err = code.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(code)
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
   AccountAuth{},
   AccountCode{},
   AccountToken{},
   HomeScreen{},
   Namer{},
   Playback{},
}
