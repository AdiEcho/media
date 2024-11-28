package kanopy

import (
   "os"
   "strings"
   "testing"
)

func TestLogin(t *testing.T) {
   email, password, ok := strings.Cut(os.Getenv("kanopy"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   data, err := web_token{}.marshal(email, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", data, os.ModePerm)
}
