package criterion

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "reflect"
   "testing"
)

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
   var pssh widevine.Pssh
   pssh.KeyId, err = hex.DecodeString(video_test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      t.Fatal(err)
   }
   var token AuthToken
   token.Raw, err = os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err = token.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   item, err := token.Video(video_test.slug)
   if err != nil {
      t.Fatal(err)
   }
   files, err := token.Files(item)
   if err != nil {
      t.Fatal(err)
   }
   file, ok := files.Dash()
   if !ok {
      t.Fatal(files.DashError())
   }
   key, err := module.Key(file, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
