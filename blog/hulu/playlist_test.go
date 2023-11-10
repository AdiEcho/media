package hulu

import (
   "154.pages.dev/http"
   "os"
   "testing"
)

func Test_Playlist(t *testing.T) {
   m, err := user_info()
   if err != nil {
      t.Fatal(err)
   }
   http.No_Location()
   http.Verbose()
   auth, err := living_room(m["username"], m["password"])
   if err != nil {
      t.Fatal(err)
   }
   res, err := auth.playlist(watch)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
