package peacock

import (
   "fmt"
   "testing"
)

// peacocktv.com/watch/playback/vod/GMO_00000000224510_02_HDSDR
const content_id = "GMO_00000000224510_02_HDSDR"

func TestPeacock(t *testing.T) {
   var video video_playouts
   err := video.New(content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
}
