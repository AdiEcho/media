package pluto

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestClip(t *testing.T) {
   for _, test := range video_tests {
      if test.clips != "" {
         clip, err := new_clip(test.clips)
         if err != nil {
            t.Fatal(err)
         }
         manifest, ok := clip.dash()
         if !ok {
            t.Fatal("episode_clip.dash")
         }
         url, err := manifest.parse(bases[0])
         if err != nil {
            t.Fatal(err)
         }
         fmt.Println(url)
         time.Sleep(time.Second)
      }
   }
}

const default_kid = "0000000063c99438d2d611a908ea7039"

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
