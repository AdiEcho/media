package ctv

import (
   "fmt"
   "testing"
   "time"
)

var test_paths = []string{
   // ctv.ca/shows/friends/the-one-with-the-bullies-s2e21
   "/shows/friends/the-one-with-the-bullies-s2e21",
   // ctv.ca/movies/the-girl-with-the-dragon-tattoo-2011
   "/movies/the-girl-with-the-dragon-tattoo-2011",
}

func TestResolvePath(t *testing.T) {
   for _, path := range test_paths {
      var resolve resolve_path
      err := resolve.New(path)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", resolve)
      time.Sleep(time.Second)
   }
}
