package criterion

import (
   "os"
   "testing"
)

func TestDelivery(t *testing.T) {
   var (
      token auth_token
      err error
   )
   token.data, err = os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   err = token.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   res, err := token.delivery()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
