package hulu

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

func TestLicense(t *testing.T) {
   var (
      protect widevine.PSSH
      err error
   )
   protect.Key_ID, err = hex.DecodeString(default_kid)
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
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   module, err := protect.CDM(private_key, client_id)
   if err != nil {
      t.Fatal(err)
   }
   auth, err := new_auth()
   if err != nil {
      t.Fatal(err)
   }
   play, err := auth.Playlist(test_deep)
   if err != nil {
      t.Fatal(err)
   }
   license, err := module.License(play)
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Printf("%x %v\n", key, ok)
}

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
const default_kid = "21b82dc2ebb24d5aa9f8631f04726650"
