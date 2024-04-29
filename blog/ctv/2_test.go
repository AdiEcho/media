package ctv

import (
   "fmt"
   "testing"
   "time"
)

func TestAxisContent(t *testing.T) {
   for _, path := range test_paths {
      resolve, err := new_resolve(path)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      axis, err := resolve.axis()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", axis)
      time.Sleep(time.Second)
   }
}
