package rakuten

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

// rakuten.tv/se/movies/i-heart-huckabees
const default_kid = "9a534a1f12d68e1a2359f38710fddb65" 

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
   var video on_demand
   video.New(classification["se"], "i-heart-huckabees")
   stream, err := video.stream()
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(stream, key_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestSe(t *testing.T) {
   var video on_demand
   video.New(classification["se"], "i-heart-huckabees")
   stream, err := video.stream()
   if err != nil {
      t.Fatal(err)
   }
   if stream.VideoQuality != "FHD" {
      t.Fatal(stream)
   }
   fmt.Printf("%+v\n", stream)
}

func TestFr(t *testing.T) {
   var video on_demand
   video.New(classification["fr"], "jerry-maguire")
   stream, err := video.stream()
   if err != nil {
      t.Fatal(err)
   }
   if stream.VideoQuality != "FHD" {
      t.Fatal(stream)
   }
}
