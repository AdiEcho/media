package rtbf

import (
   "fmt"
   "os"
   "testing"
)

func TestSix(t *testing.T) {
   var o one
   err := o.New()
   if err != nil {
      t.Fatal(err)
   }
   text, err := os.ReadFile("account.json")
   if err != nil {
      t.Fatal(err)
   }
   var account accounts_login
   err = account.unmarshal(text)
   if err != nil {
      t.Fatal(err)
   }
   token, err := o.four(&account)
   if err != nil {
      t.Fatal(err)
   }
   gigya, err := token.login()
   if err != nil {
      t.Fatal(err)
   }
   title, err := gigya.entitlement()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", title)
}
