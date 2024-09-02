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
      clip, err := OnDemand{Id: test.id}.Clip()
      if err != nil {
         t.Fatal(err)
      }
      manifest, ok := clip.Dash()
      if !ok {
         t.Fatal("EpisodeClip.Dash")
      }
      manifest.Url.Scheme = Bases[0].Scheme
      manifest.Url.Host = Bases[0].Host
      fmt.Printf("%+v\n", manifest)
      time.Sleep(time.Second)
   }
}

// the slug is useful as it sometimes contains the year, but its not worth
// parsing since its sometimes missing
var video_tests = []struct{
   id string
   key_id string
   url   string
}{
   {
      id: "5c4bb2b308d10f9a25bbc6af",
      key_id: "0000000066bfe3cd26602c92dc082e3b",
      url: "pluto.tv/on-demand/movies/bound-paramount-1-1",
   },
   {
      id: "66b3838317101c00130b411e",
      key_id: "0000000066b3c161c1cee84ffce71de3",
      url: "pluto.tv/on-demand/movies/just-go-with-it-2011-1-1",
   },
   {
      id: "6356d14136d64a001450b121",
      key_id: "000000006358c035248b647dad3c09ad",
      url: "pluto.tv/on-demand/series/frasier-cbs-tv/season/1/episode/space-quest-1992-1-2",
   },
}

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
   for _, test := range video_tests {
      var pssh widevine.Pssh
      pssh.KeyId, err = hex.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(Poster{}, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}
