package rtbf

import (
   "os"
   "testing"
)

func TestFour(t *testing.T) {
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
   res, err := o.four(login)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
