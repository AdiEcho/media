package ctv

import (
   "154.pages.dev/encoding"
   "fmt"
   "testing"
   "time"
)

func TestMedia(t *testing.T) {
   for _, path := range test_paths {
      resolve, err := new_resolve(path)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
      axis, err := resolve.axis()
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
      media, err := axis.media()
      if err != nil {
         t.Fatal(err)
      }
      name, err := encoding.Name(namer{media})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(99 * time.Millisecond)
   }
}
