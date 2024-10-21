package max

import (
   "fmt"
   "testing"
)

func TestOne(t *testing.T) {
   var token bolt_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(token.st)
}
