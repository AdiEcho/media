package web

import (
   "154.pages.dev/media/blog/spotify/android"
   "fmt"
   "os"
   "testing"
)

const file_id = "392482fe9bed7372d1657d7e22f32b792902f3bd"

func TestStorage(t *testing.T) {
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
   var storage storage_resolve
   if err := storage.New(login, file_id); err != nil {
      t.Fatal(err)
   }
   for _, url := range storage.CDNURL {
      fmt.Println(url)
   }
}
