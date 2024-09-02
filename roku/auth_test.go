package roku

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestContent(t *testing.T) {
   for _, test := range tests {
      var home HomeScreen
      err := home.New(test.id)
      if err != nil {
         t.Fatal(err)
      }
      name, err := text.Name(Namer{home})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}

var tests = map[string]struct {
   id string
   key_id string
   url string
} {
   "episode": {
      id: "105c41ea75775968b670fbb26978ed76",
      key_id: "bdfa4d6cdb39702e5b681f90617f9a7e",
      url: "therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76",
   },
   "movie": {
      id: "597a64a4a25c5bf6af4a8c7053049a6f",
      key_id: "28339ad78f734520da24e6e0573d392e",
      url: "therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f",
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
   for _, test := range tests{
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
      var auth AccountAuth
      err = auth.New(nil)
      if err != nil {
         t.Fatal(err)
      }
      err = auth.Unmarshal()
      if err != nil {
         t.Fatal(err)
      }
      play, err := auth.Playback(test.id)
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(play, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}
