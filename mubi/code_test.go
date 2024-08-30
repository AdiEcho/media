package mubi

import (
   "fmt"
   "os"
   "testing"
)

func TestCode(t *testing.T) {
   var code LinkCode
   err := code.New()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.txt", code.Raw, os.ModePerm)
   err = code.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(code)
}
