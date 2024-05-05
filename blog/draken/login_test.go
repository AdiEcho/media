package draken

import (
   "fmt"
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
   fmt.Printf("%+v\n", auth)
}
