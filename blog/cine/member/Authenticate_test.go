package member

import (
   "os"
   "testing"
)

func TestAuthenticate(t *testing.T) {
   username := os.Getenv("cineMember_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("cineMember_password")
   var auth authenticate
   err := auth.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.json", auth.data, 0666)
}
