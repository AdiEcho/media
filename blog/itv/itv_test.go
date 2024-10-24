package itv

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "fmt"
   "os"
   "testing"
   "time"
)

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

var tests = []struct{
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
func TestPlaylist(t *testing.T) {
   var play playlist
   err := play.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(play.resolution_720())
}

const (
   content_id = "10-3918-0001-001_34"
   key_id = "\xcex\xf3*\x03\x9aD\x1a\x890\x10Ɖ@\xcd\xf2"
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
   pssh.ContentId = []byte(content_id)
   pssh.KeyId = []byte(key_id)
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(poster{}, []byte(key_id))
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
