package member

import (
   "os"
   "strings"
   "testing"
)

func TestAuthenticate(t *testing.T) {
   username, password, ok := strings.Cut(os.Getenv("cine_member"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   var user OperationUser
   err := user.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("user.txt", user.Raw, os.ModePerm)
}
