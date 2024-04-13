package tubi

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

func TestLicense(t *testing.T) {
   test := tests["movie"]
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
   key_id, err := hex.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   if err := module.New(private_key, client_id, key_id); err != nil {
      t.Fatal(err)
   }
   var cms content
   if err := cms.New(test.content_id); err != nil {
      t.Fatal(err)
   }
   video, ok := cms.resolution_720p()
   if !ok {
      t.Fatal("resolution_720p")
   }
   license, err := module.License(video)
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Printf("%x %v\n", key, ok)
}
