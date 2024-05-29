package roku

import (
   "fmt"
   "os"
   "testing"
)

func TestTwo(t *testing.T) {
   var one one_response
   err := one.New()
   if err != nil {
      t.Fatal(err)
   }
   two, err := one.two()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(two)
   text, err := two.marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("2.json", text, 0666)
}
