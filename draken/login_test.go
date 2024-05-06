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
   var auth auth_login
   err := auth.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("login.json", auth.data, 0666)
}
