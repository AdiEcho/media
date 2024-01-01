package youtube

import (
   "fmt"
   "testing"
   "time"
)

func Test_Code(t *testing.T) {
   var code Device_Code
   err := code.Post()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf(
      "1. go to\n%v\n\n2. enter this code\n%v\n",
      code.Verification_URL, code.User_Code,
   )
   for range [9]bool{} {
      time.Sleep(9 * time.Second)
      raw, err := code.Token()
      if err != nil {
         t.Fatal(err)
      }
      var tok Token
      tok.Unmarshal(raw)
      fmt.Printf("%+v\n", tok)
      if tok.Access_Token != "" {
         break
      }
   }
}
