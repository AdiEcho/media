package peacock

import (
   "fmt"
   "os"
   "testing"
)

func TestSignin(t *testing.T) {
   user, password := os.Getenv("peacock_username"), os.Getenv("peacock_password")
   if user == "" {
      t.Fatal("peacock_username")
   }
   cookies, err := signin(user, password)
   if err != nil {
      t.Fatal(err)
   }
   for i, cookie := range cookies {
      if i >= 1 {
         fmt.Println()
      }
      fmt.Println(cookie)
   }
}
