package stan

import (
   "fmt"
   "os"
   "testing"
)

func TestCode(t *testing.T) {
   var code activation_code
   err := code.New()
   if err != nil {
      t.Fatal(err)
   }
   code.unmarshal()
   fmt.Println(code)
   os.WriteFile("1.json", code.data, 0666)
}
