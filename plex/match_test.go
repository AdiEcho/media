package plex

import (
   "os"
   "testing"
   "time"
)

var paths = []string{
   //"/movie/cruel-intentions",
   "/show/broadchurch/season/3/episode/5",
}

func TestMatch(t *testing.T) {
   var anon anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, path := range paths {
      func() {
         res, err := anon.matches(path)
         if err != nil {
            t.Fatal(err)
         }
         defer res.Body.Close()
         res.Write(os.Stdout)
      }()
      time.Sleep(time.Second)
   }
}
