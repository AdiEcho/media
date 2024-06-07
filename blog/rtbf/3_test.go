package rtbf

import (
   "os"
   "testing"
)

func TestAccountsLogin(t *testing.T) {
   var o one
   err := o.New()
   if err != nil {
      t.Fatal(err)
   }
   username, password := os.Getenv("rtbf_username"), os.Getenv("rtbf_password")
   if username == "" {
      t.Fatal("Getenv")
   }
   login, err := o.login(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("account.json", login.data, 0666)
}
