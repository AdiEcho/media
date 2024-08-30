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
   os.WriteFile("code.txt", code.Data, os.ModePerm)
   code.Unmarshal()
   fmt.Println(code)
}
