package cbc

import (
   "os"
   "testing"
)

func TestProfile(t *testing.T) {
   username, password := os.Getenv("cbc_username"), os.Getenv("cbc_password")
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var token LoginToken
   if err := token.New(username, password); err != nil {
      t.Fatal(err)
   }
   pro, err := token.Profile()
   if err != nil {
      t.Fatal(err)
   }
   if err := pro.WriteFile(home + "/cbc/profile.json"); err != nil {
      t.Fatal(err)
   }
}
