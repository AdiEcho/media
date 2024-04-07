package plex

import (
   "fmt"
   "testing"
)

const (
   // watch.plex.tv/movie/cruel-intentions
   movie = "/movie/cruel-intentions"
   default_kid = "eabdd790d9279b9699b32110eed9a154"
   episode = "/show/broadchurch/season/3/episode/5"
)

func TestDiscover(t *testing.T) {
   var anon anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   meta, err := anon.discover(movie)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(meta)
}

