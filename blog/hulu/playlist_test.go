package hulu

import (
   "154.pages.dev/http"
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
)

const raw_pssh = ""

func Test_License(t *testing.T) {
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
   pssh, err := base64.StdEncoding.DecodeString(raw_pssh)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, pssh)
   if err != nil {
      t.Fatal(err)
   }
   http.No_Location()
   http.Verbose()
   play, err := new_playlist()
   if err != nil {
      t.Fatal(err)
   }
   key, err := mod.Key(play)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
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

func Test_Playlist(t *testing.T) {
   http.No_Location()
   http.Verbose()
   play, err := new_playlist()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", play)
}
