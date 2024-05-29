package roku

import (
   "fmt"
   "os"
   "testing"
)

func TestAccountToken(t *testing.T) {
   var (
      activate activation_token
      err error
   )
   activate.data, err = os.ReadFile("3.json")
   if err != nil {
      t.Fatal(err)
   }
   err = activate.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   var account account_token
   err = account.New(&activate)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", account)
}
