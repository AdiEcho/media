package peacock

import (
   "fmt"
   "os"
   "testing"
)

// peacocktv.com/watch/playback/vod/GMO_00000000224510_02_HDSDR
const content_id = "GMO_00000000224510_02_HDSDR"

func TestVideo(t *testing.T) {
   user, password := os.Getenv("peacock_username"), os.Getenv("peacock_password")
   if user == "" {
      t.Fatal("peacock_username")
   }
   var sign sign_in
   err := sign.New(user, password)
   if err != nil {
      t.Fatal(err)
   }
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
