package web

import (
   "154.pages.dev/media/blog/spotify/android"
   "fmt"
   "os"
   "testing"
)

// open.spotify.com/track/1oaaSrDJimABpOdCEbw2DJ
const track = "2da9a11032664413b24de181c534f157"

func TestMetadata(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var login android.LoginOk
   login.Data, err = os.ReadFile(home + "/spotify.bin")
   if err != nil {
      t.Fatal(err)
   }
   if err := login.Consume(); err != nil {
      t.Fatal(err)
   }
   var meta metadata
   if err := meta.New(login, track); err != nil {
      t.Fatal(err)
   }
   for _, file := range meta.File {
      fmt.Printf("%+v\n", file)
   }
}
