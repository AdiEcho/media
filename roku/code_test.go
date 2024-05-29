package roku

import (
   "fmt"
   "os"
   "testing"
)

func TestActivationCode(t *testing.T) {
   var token account_token
   err := token.New(nil)
   if err != nil {
      t.Fatal(err)
   }
   code, err := token.code()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(code)
   text, err := code.marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.json", text, 0666)
}
