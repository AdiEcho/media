package rtbf

import (
   "os"
   "testing"
)

func TestAccountsLogin(t *testing.T) {
   username := os.Getenv("rtbf_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("rtbf_password")
   var login AuvioLogin
   err := login.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("login.txt", login.Raw, os.ModePerm)
}
