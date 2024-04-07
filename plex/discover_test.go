package plex

import (
   "154.pages.dev/encoding"
   "fmt"
   "testing"
)

var tests = map[string]string{
   "episode": "/show/broadchurch/season/3/episode/5",
   // watch.plex.tv/movie/cruel-intentions
   "movie": "/movie/cruel-intentions",
}

func TestDiscover(t *testing.T) {
   var anon anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      match, err := anon.discover(test)
      if err != nil {
         t.Fatal(err)
      }
      name, err := encoding.Name(match)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
}

