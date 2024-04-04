package roku

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestLicense(t *testing.T) {
   var site CrossSite
   if err := site.New(); err != nil {
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
   for _, test := range tests {
      var protect widevine.PSSH
      {
         b, err := base64.StdEncoding.DecodeString(test.pssh)
         if err != nil {
            t.Fatal(err)
         }
         if err := protect.New(b); err != nil {
            t.Fatal(err)
         }
      }
      module, err := protect.CDM(private_key, client_id)
      if err != nil {
         t.Fatal(err)
      }
      play, err := site.Playback(test.playback_id)
      if err != nil {
         t.Fatal(err)
      }
      license, err := module.License(play)
      if err != nil {
         t.Fatal(err)
      }
      key, ok := module.Key(license)
      fmt.Println(key, ok)
   }
}

func TestPlayback(t *testing.T) {
   var site CrossSite
   err := site.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      play, err := site.Playback(test.playback_id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(play)
      time.Sleep(time.Second)
   }
}
