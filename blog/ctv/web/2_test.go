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
      content, err := resolve.content()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", content)
      time.Sleep(time.Second)
   }
}
