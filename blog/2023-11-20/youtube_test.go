package youtube

import (
   "fmt"
   "testing"
   "time"
)

var ids = []string{
   "2ZcDwdXEVyI", // episode
   "7KLCti7tOXE", // video
   "zZUUHthMeA4", // film
}

func Test_Watch(t *testing.T) {
   for _, id := range ids {
      c, err := make_contents(id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", c)
      time.Sleep(time.Second)
   }
}
