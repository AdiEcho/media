package criterion

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

// criterionchannel.com/videos/my-dinner-with-andre
const default_kid = "e4576465a745213f336c1ef1bf5d513e"

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
   err = module.New(private_key, client_id, widevine.PSSH(key_id, nil))
   if err != nil {
      t.Fatal(err)
   }
   var token AuthToken
   token.data, err = os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   token.unmarshal()
   item, err := token.video(my_dinner)
   if err != nil {
      t.Fatal(err)
   }
   deliver, err := token.delivery(item)
   if err != nil {
      t.Fatal(err)
   }
   stream, ok := deliver.dash()
   if !ok {
      t.Fatal("video_delivery.dash")
   }
   key, err := module.Key(stream, key_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
