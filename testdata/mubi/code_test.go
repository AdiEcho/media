package mubi

import (
   "fmt"
   "os"
   "testing"
)

func TestCode(t *testing.T) {
   var code linkCode
   err := code.New("US")
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.json", code.Raw, 0666)
   code.unmarshal()
   fmt.Println(code)
}
