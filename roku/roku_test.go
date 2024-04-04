package roku

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "path"
   "testing"
   "time"
)

func TestPlayback(t *testing.T) {
   var site CrossSite
   err := site.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      play, err := site.Playback(path.Base(test.url))
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(play)
      time.Sleep(time.Second)
   }
}

func TestLicense(t *testing.T) {
   test := tests["episode"]
   var site CrossSite
   if err := site.New(); err != nil {
      t.Fatal(err)
   }
   play, err := site.Playback(path.Base(test.url))
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
   key_id, err := hex.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   if err := module.New(private_key, client_id, key_id); err != nil {
      t.Fatal(err)
   }
   license, err := module.License(play)
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Println(key, ok)
}
