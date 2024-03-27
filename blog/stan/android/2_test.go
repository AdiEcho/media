package stan

import (
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
   os.WriteFile("2.json", token.data, 0666)
}
