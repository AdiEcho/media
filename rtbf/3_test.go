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
   var login accounts_login
   err := login.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   text, err := login.marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("account.json", text, 0666)
}
