package roku

import (
   "fmt"
   "testing"
)

func TestContent(t *testing.T) {
   u, err := content("HELLO")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(u)
}
