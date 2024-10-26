package max

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestInitiate(t *testing.T) {
   var token BoltToken
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", []byte(token.St), os.ModePerm)
   time.Sleep(time.Second)
   initiate, err := token.Initiate()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", initiate)
}
