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
   os.WriteFile("code.json", code.data, 0666)
}

func TestToken(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var code activation_code
   code.data, err = os.ReadFile("code.json")
   if err != nil {
      t.Fatal(err)
   }
   code.unmarshal()
   token, err := code.token()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile(home + "/stan.json", token.data, 0666)
}
