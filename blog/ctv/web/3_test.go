package ctv

import (
   "os"
   "testing"
   "time"
)

func TestContentPackages(t *testing.T) {
   for _, path := range test_paths {
      resolve, err := new_resolve(path)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
      content, err := resolve.content()
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
      func() {
         res, err := content.content_packages()
         if err != nil {
            t.Fatal(err)
         }
         defer res.Body.Close()
         res.Write(os.Stdout)
      }()
      break
   }
}
