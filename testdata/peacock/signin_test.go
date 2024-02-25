package peacock

import (
   "fmt"
   "os"
   "testing"
)

func TestTokens(t *testing.T) {
   user, password := os.Getenv("peacock_username"), os.Getenv("peacock_password")
   if user == "" {
      t.Fatal("peacock_username")
   }
   var sign sign_in
   err := sign.New(user, password)
   if err != nil {
      t.Fatal(err)
   }
   auth, err := sign.auth()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", auth)
}

func TestSignin(t *testing.T) {
   user, password := os.Getenv("peacock_username"), os.Getenv("peacock_password")
   if user == "" {
      t.Fatal("peacock_username")
   }
   var sign sign_in
   err := sign.New(user, password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(sign.cookie)
}
