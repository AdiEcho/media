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
   os.WriteFile("auth.txt", auth.Raw, os.ModePerm)
   err = auth.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   // AccountCode
   code, err := auth.Code()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.txt", code.Raw, os.ModePerm)
   err = code.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(code)
}
