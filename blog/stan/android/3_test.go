package stan

import (
   "os"
   "testing"
)

func TestSession(t *testing.T) {
   var (
      code activation_code
      err error
   )
   code.data, err = os.ReadFile("1.json")
   if err != nil {
      t.Fatal(err)
   }
   code.unmarshal()
   token, err := code.token()
   if err != nil {
      t.Fatal(err)
   }
   res, err := token.session()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
