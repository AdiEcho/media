package rtbf

import (
   "os"
   "testing"
)

func TestSix(t *testing.T) {
   var o one
   err := o.New()
   if err != nil {
      t.Fatal(err)
   }
   var account accounts_login
   account.data, err = os.ReadFile("account.json")
   if err != nil {
      t.Fatal(err)
   }
   err = account.unmarshal()
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
   res, err := gigya.six()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
