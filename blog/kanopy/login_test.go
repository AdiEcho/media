package kanopy

import (
   "fmt"
   "os"
   "strings"
   "testing"
)

func TestLogin(t *testing.T) {
   email, password, ok := strings.Cut(os.Getenv("kanopy"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   var web web_token
   err := web.New(email, password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", web)
}
