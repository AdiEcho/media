package roku

import (
   "fmt"
   "os"
   "testing"
)

func TestActivationCode(t *testing.T) {
   var token AccountToken
   err := token.New(nil)
   if err != nil {
      t.Fatal(err)
   }
   code, err := token.Code()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(code)
   text, err := code.Marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.json", text, 0666)
}
