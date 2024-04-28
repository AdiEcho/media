package ctv

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

const default_kid = "cb09571eebcb3f7287202657f6b9f7a6"

func TestLicense(t *testing.T) {
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
   key_id, err := hex.DecodeString(default_kid)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   if err := module.New(private_key, client_id, key_id); err != nil {
      t.Fatal(err)
   }
   license, err := module.License(poster{})
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Printf("%x %v\n", key, ok)
}
