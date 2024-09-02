package roku

import (
   "fmt"
   "os"
   "testing"
)

func TestCode(t *testing.T) {
   // AccountAuth
   var auth AccountAuth
   err := auth.New(nil)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("auth.txt", auth.Data, os.ModePerm)
   auth.Unmarshal()
   // AccountCode
   code, err := auth.Code()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.txt", code.Data, os.ModePerm)
   code.Unmarshal()
   fmt.Println(code)
}
