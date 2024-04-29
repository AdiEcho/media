package ctv

import (
   "fmt"
   "testing"
   "time"
)

func TestResolvePath(t *testing.T) {
   for _, path := range test_paths {
      resolve, err := new_resolve(path)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(resolve.id())
      time.Sleep(time.Second)
   }
}
