package hulu

import (
   "154.pages.dev/http"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

func Test_Playlist(t *testing.T) {
   http.No_Location()
   http.Verbose()
   play, err := new_playlist()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", play)
}

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
const default_KID = "21b82dc2ebb24d5aa9f8631f04726650"

func new_playlist() (*Playlist, error) {
   m, err := user_info()
   if err != nil {
      return nil, err
   }
   auth, err := Living_Room(m["username"], m["password"])
   if err != nil {
      return nil, err
   }
   auth.Unmarshal()
   return auth.Playlist(test_deep)
}

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
   kid, err := hex.DecodeString(default_KID)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, kid, nil)
   if err != nil {
      t.Fatal(err)
   }
   key, err := mod.Key(play)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
