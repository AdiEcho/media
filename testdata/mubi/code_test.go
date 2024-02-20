package mubi

import (
   "fmt"
   "testing"
)

func TestCode(t *testing.T) {
   var code linkCode
   err := code.New("US")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", code)
}
