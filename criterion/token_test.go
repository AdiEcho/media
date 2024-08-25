package criterion

import (
   "os"
   "testing"
)

func TestToken(t *testing.T) {
   username := os.Getenv("criterion_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("criterion_password")
   var token AuthToken
   err := token.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", token.Raw, 0666)
}
