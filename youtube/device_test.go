package youtube

import (
   "fmt"
   "testing"
   "time"
)

func Test_Code(t *testing.T) {
   code, err := New_Device_Code()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf(
      "1. go to\n%v\n\n2. enter this code\n%v\n",
      code.Verification_URL, code.User_Code,
   )
   for range [9]bool{} {
      time.Sleep(9 * time.Second)
      tok, err := code.Token()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", tok)
      if tok.Access_Token != "" {
         break
      }
   }
}
