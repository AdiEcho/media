package rtbf

import (
   "os"
   "testing"
)

func TestAccountsLogin(t *testing.T) {
   username, password := os.Getenv("rtbf_username"), os.Getenv("rtbf_password")
   if username == "" {
      t.Fatal("Getenv")
   }
   var o one
   err := o.New()
   if err != nil {
      t.Fatal(err)
   }
   login, err := o.login(username, password)
   if err != nil {
      t.Fatal(err)
   }
   text, err := login.marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("account.json", text, 0666)
}
