package plex

import (
   "154.pages.dev/text"
   "fmt"
   "testing"
   "time"
)

var tests = map[string]Path{
   "episode": {"/show/broadchurch/season/3/episode/5"},
   // watch.plex.tv/movie/cruel-intentions
   "movie": {"/movie/cruel-intentions"},
}

func TestDiscover(t *testing.T) {
   var anon Anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      match, err := anon.Discover(test)
      if err != nil {
         t.Fatal(err)
      }
      name, err := text.Name(Namer{match})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
}
