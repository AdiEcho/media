package paramount

import (
   "fmt"
   "testing"
)

func TestMpdFr(t *testing.T) {
   var head Header
   err := head.New(tests["fr"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", head)
}

func TestMpdUs(t *testing.T) {
   var head Header
   err := head.New(tests["us"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", head)
}
