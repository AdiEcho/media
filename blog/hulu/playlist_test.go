package hulu

import (
   "154.pages.dev/http"
   "154.pages.dev/widevine"
   "fmt"
   "os"
   "testing"
)

func Test_License(t *testing.T) {
   http.No_Location()
   http.Verbose()
   play, err := new_playlist()
   if err != nil {
      t.Fatal(err)
   }
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   // THIS IS NOT GOOD
   mod, err := widevine.New_Module(private_key, client_ID, nil)
   if err != nil {
      t.Fatal(err)
   }
   key, err := mod.Key(play)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func Test_Playlist(t *testing.T) {
   http.No_Location()
   http.Verbose()
   play, err := new_playlist()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", play)
}

func new_playlist() (*playlist, error) {
   m, err := user_info()
   if err != nil {
      return nil, err
   }
   auth, err := living_room(m["username"], m["password"])
   if err != nil {
      return nil, err
   }
   return auth.playlist(watch)
}
