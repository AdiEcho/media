package joyn

import (
   "fmt"
   "testing"
)

// joyn.de/serien/one-tree-hill/1-2-quaelende-angst
const one_tree = "/serien/one-tree-hill/1-2-quaelende-angst"

func TestEpisode(t *testing.T) {
   var episode episode_detail
   //err := episode.New(one_tree)
   err := episode.New(barry_seal)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", episode)
}
