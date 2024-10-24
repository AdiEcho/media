package itv

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = []struct{
   content_id string
   key_id string
   legacy_id string
   url string
}{
   {
      legacy_id: "10/3463/0001",
      url: "itv.com/watch/pulp-fiction/10a3463",
   },
   {
      legacy_id: "10/3915/0002",
      url: "itv.com/watch/community/10a3915/10a3915a0002",
   },
}

func TestDiscovery(t *testing.T) {
   for _, test := range tests {
      var title discovery_title
      err := title.New(test.legacy_id)
      if err != nil {
         t.Fatal(err)
      }
      name, err := text.Name(namer{title})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}

func TestPlaylist(t *testing.T) {
   for _, test := range tests {
      var title discovery_title
      err := title.New(test.legacy_id)
      if err != nil {
         t.Fatal(err)
      }
      play, err := title.playlist()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(play.resolution_720())
      time.Sleep(time.Second)
   }
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
   for _, test := range tests {
      var pssh widevine.Pssh
      pssh.ContentId = []byte(test.content_id)
      pssh.KeyId = []byte(test.key_id)
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(poster{}, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}
