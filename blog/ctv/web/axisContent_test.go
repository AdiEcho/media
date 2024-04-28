package ctv

import (
   "fmt"
   "testing"
   "time"
)

var axis_ids = []string{
   // ctv.ca/movies/the-girl-with-the-dragon-tattoo-2011
   "contentid/axis-content-1417780",
   // ctv.ca/shows/friends/the-one-with-the-bullies-s2e21
   "contentid/axis-content-1730820",
}

func TestAxisContent(t *testing.T) {
   for _, id := range axis_ids {
      var axis axis_content
      err := axis.New(id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", axis)
      time.Sleep(time.Second)
   }
}
