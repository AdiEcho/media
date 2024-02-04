package hulu

import (
   "154.pages.dev/widevine"
   "fmt"
   "os"
   "testing"
)

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
const default_kid = "21b82dc2ebb24d5aa9f8631f04726650"

func Test_License(t *testing.T) {
   var module widevine.CDM
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   if err := module.New(private_key); err != nil {
      t.Fatal(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   kid, err := hex.DecodeString(default_kid)
   if err != nil {
      t.Fatal(err)
   }
   module.Key_ID(client_id, kid)
   play, err := new_playlist()
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(play)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

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
