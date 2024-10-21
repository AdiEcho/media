package max

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestInitiate(t *testing.T) {
   var token bolt_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", []byte(token.st), os.ModePerm)
   ///////////////////////
   time.Sleep(time.Second)
   initiate, err := token.initiate()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", initiate)
}
