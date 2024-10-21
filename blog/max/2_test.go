package max

import (
   "os"
   "testing"
)

func TestTwo(t *testing.T) {
   var token bolt_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   resp, err := token.initiate()
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
