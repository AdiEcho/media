package max

import (
   "fmt"
   "testing"
)

func TestDefaultToken(t *testing.T) {
   var token default_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", token)
}
