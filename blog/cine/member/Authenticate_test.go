package member

import (
   "fmt"
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
   fmt.Printf("%+v\n", auth)
}
