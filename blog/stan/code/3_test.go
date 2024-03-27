package stan

import (
   "fmt"
   "os"
   "testing"
)

func TestSession(t *testing.T) {
   var (
      token web_token
      err error
   )
   token.data, err = os.ReadFile("2.json")
   if err != nil {
      t.Fatal(err)
   }
   token.unmarshal()
   session, err := token.session()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", session)
}
