package rtbf

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
)

func TestSeven(t *testing.T) {
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
   content_id, err := base64.StdEncoding.DecodeString(raw_content_id)
   if err != nil {
      t.Fatal(err)
   }
   key_id, err := base64.StdEncoding.DecodeString(raw_key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, widevine.PSSH(key_id, content_id))
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(poster{}, key_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

const (
   raw_content_id = "bzFDMzdUdDVTem1ITW1FZ1FWaVVFQT09"
   raw_key_id = "o1C37Tt5SzmHMmEgQViUEA=="
)
