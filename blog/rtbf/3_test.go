package rtbf

import (
   "fmt"
   "os"
   "testing"
)

func TestTwo(t *testing.T) {
   var o one
   err := o.New()
   if err != nil {
      t.Fatal(err)
   }
   username, password := os.Getenv("rtbf_username"), os.Getenv("rtbf_password")
   if username == "" {
      t.Fatal("Getenv")
   }
   login, err := o.login(username, password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", login)
}
