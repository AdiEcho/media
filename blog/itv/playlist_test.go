package itv

import (
   "fmt"
   "testing"
)

func TestPlaylist(t *testing.T) {
   var play playlist
   err := play.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(play.resolution_720())
}
