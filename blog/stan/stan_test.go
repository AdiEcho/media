package stan

import (
   "fmt"
   "os"
   "testing"
)

func TestToken(t *testing.T) {
   var (
      code activation_code
      err error
   )
   code.data, err = os.ReadFile("1.json")
   if err != nil {
      t.Fatal(err)
   }
   code.unmarshal()
   token, err := code.token()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.json", token.data, 0666)
}

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
