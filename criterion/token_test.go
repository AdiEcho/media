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
   data, err := NewAuthToken(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.json", data, 0666)
}
