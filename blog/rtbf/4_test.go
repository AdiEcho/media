package rtbf

import (
   "fmt"
   "os"
   "testing"
)

func TestFour(t *testing.T) {
   var o one
   err := o.New()
   if err != nil {
      t.Fatal(err)
   }
   var login accounts_login
   login.data, err = os.ReadFile("login.json")
   if err != nil {
      t.Fatal(err)
   }
   err = login.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   token, err := o.four(&login)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", token)
}
