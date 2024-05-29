package roku

import (
   "os"
   "testing"
)

func TestActivationToken(t *testing.T) {
   text, err := os.ReadFile("2.json")
   if err != nil {
      t.Fatal(err)
   }
   var code activation_code
   err = code.unmarshal(text)
   if err != nil {
      t.Fatal(err)
   }
   token, err := code.token()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("3.json", token.data, 0666)
}
