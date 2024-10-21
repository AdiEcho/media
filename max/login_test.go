package max

import (
   "fmt"
   "os"
   "testing"
)

func TestLogin(t *testing.T) {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token bolt_token
   token.st = string(data)
   login, err := token.login()
   if err != nil {
      t.Fatal(err)
   }
   err = login.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", login.Data)
}
