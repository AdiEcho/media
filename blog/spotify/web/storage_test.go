package web

import (
   "154.pages.dev/media/blog/spotify/android"
   "encoding/json"
   "fmt"
   "os"
   "slices"
   "testing"
   "time"
)

func TestStorage(t *testing.T) {
   text, err := os.ReadFile("metadata.json")
   if err != nil {
      t.Fatal(err)
   }
   var meta metadata
   if err := json.Unmarshal(text, &meta); err != nil {
      t.Fatal(err)
   }
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
   for _, file := range meta.File {
      fmt.Println(file.Format)
      var storage storage_resolve
      if err := storage.New(login, file.File_ID); err != nil {
         t.Fatal(err)
      }
      slices.SortFunc(storage.CDNURL, func(a, b string) int {
         return len(a) - len(b)
      })
      fmt.Println(storage.CDNURL[0])
      time.Sleep(time.Second)
   }
}
