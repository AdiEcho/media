package roku

import (
   "fmt"
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
}
