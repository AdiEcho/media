package draken

import (
   "os"
   "testing"
)

func TestLogin(t *testing.T) {
   username := os.Getenv("draken_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("draken_password")
   var login AuthLogin
   err := login.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("login.json", login.Data, 0666)
}
