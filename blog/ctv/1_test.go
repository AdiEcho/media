package ctv

import (
   "fmt"
   "testing"
   "time"
)

var test_paths = []string{
   // ctv.ca/movies/the-girl-with-the-dragon-tattoo-2011
   "/movies/the-girl-with-the-dragon-tattoo-2011",
   // ctv.ca/shows/friends/the-one-with-the-bullies-s2e21
   "/shows/friends/the-one-with-the-bullies-s2e21",
}

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
