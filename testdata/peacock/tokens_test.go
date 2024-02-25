package peacock

import (
   "fmt"
   "testing"
)

func TestTokens(t *testing.T) {
   var auth auth_tokens
   err := auth.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", auth)
}
