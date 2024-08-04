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
   var login AccountLogin
   err := login.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   text, err := login.Marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("account.txt", text, 0666)
}
