package stan

import (
   "fmt"
   "testing"
)

func TestCode(t *testing.T) {
   var code activation_code
   err := code.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(code)
}
