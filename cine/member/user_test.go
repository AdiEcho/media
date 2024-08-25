package member

import (
   "os"
   "testing"
)

func TestAuthenticate(t *testing.T) {
   username := os.Getenv("cine_member_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("cine_member_password")
   var user OperationUser
   err := user.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("user.txt", user.Raw, 0666)
}
