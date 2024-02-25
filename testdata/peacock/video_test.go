package peacock

import (
   "fmt"
   "os"
   "testing"
)

// peacocktv.com/watch/playback/vod/GMO_00000000224510_02_HDSDR
const content_id = "GMO_00000000224510_02_HDSDR"

func TestVideo(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   text, err := os.ReadFile(home + "/peacock.json")
   if err != nil {
      t.Fatal(err)
   }
   var sign sign_in
   sign.unmarshal(text)
   auth, err := sign.auth()
   if err != nil {
      t.Fatal(err)
   }
   video, err := auth.video(content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
}
