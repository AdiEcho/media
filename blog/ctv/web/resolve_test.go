package ctv

import (
   "fmt"
   "testing"
)

func TestResolve(t *testing.T) {
   var path resolve_path
   err := path.New("/movies/the-girl-with-the-dragon-tattoo-2011")
   if err != nil {
      t.Fatal(err)
   }
   content, err := path.content()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", content)
}
