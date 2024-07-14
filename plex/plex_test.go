package plex

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
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
   pssh.KeyId, err = hex.DecodeString(default_kid)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Encode())
   if err != nil {
      t.Fatal(err)
   }
   var anon Anonymous
   err = anon.New()
   if err != nil {
      t.Fatal(err)
   }
   match, err := anon.Discover(tests["movie"])
   if err != nil {
      t.Fatal(err)
   }
   video, err := anon.Video(match, "")
   if err != nil {
      t.Fatal(err)
   }
   part, ok := video.DASH(anon)
   if !ok {
      t.Fatal("metadata.dash")
   }
   key, err := module.Key(part, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

const default_kid = "eabdd790d9279b9699b32110eed9a154"
