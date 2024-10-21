package max

import (
   "os"
   "testing"
   "time"
)

func TestThree(t *testing.T) {
   var token bolt_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   time.Sleep(time.Second)
   resp, err := token.initiate()
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
   time.Sleep(time.Second)
   func() {
      resp, err := token.login()
      if err != nil {
         t.Fatal(err)
      }
      defer resp.Body.Close()
      resp.Write(os.Stdout)
   }()
}
