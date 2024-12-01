package plex

import (
   "41.neocities.org/media/plex"
   "41.neocities.org/widevine"
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
   var user plex.Anonymous
   err = user.New()
   if err != nil {
      t.Fatal(err)
   }
   match, err := user.Match(&plex.Address{"/movie/never-look-away-2024"})
   if err != nil {
      t.Fatal(err)
   }
   video, err := user.Video(match, "")
   if err != nil {
      t.Fatal(err)
   }
   part, ok := video.Dash()
   if !ok {
      t.Fatal("Metadata.Dash")
   }
   var pssh widevine.Pssh
   pssh.KeyId, err = hex.DecodeString("65adf42460cf5e5a20aa728f0e4b8680")
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(part, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
