package plex

import (
   "fmt"
   "testing"
)

const cruel = "https://watch.plex.tv/movie/cruel-intentions"

func TestMetadata(t *testing.T) {
   var anon anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   meta, err := anon.metadata(cruel)
   if err != nil {
      t.Fatal(err)
   }
   part, ok := meta.dash(anon.AuthToken)
   fmt.Printf("%+v %v\n", part, ok)
}
