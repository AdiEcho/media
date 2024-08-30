package nbc

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

const key_id = "0552e44842654a4e81b326004be47be0"

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
   pssh.KeyId, err = hex.DecodeString(key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(Core(), pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestVideo(t *testing.T) {
   v, ok := Core().RequestUrl()
   fmt.Println(v, ok)
}
var mpx_guids = []int{
   // episode unlocked
   // nbc.com/saturday-night-live/video/may-18-jake-gyllenhaal/9000283438
   9000283422,
   // movie locked
   // nbc.com/2-fast-2-furious/video/2-fast-2-furious/2957739
   2957739,
}

func TestMetadata(t *testing.T) {
   for _, mpx_guid := range mpx_guids {
      meta, err := NewMetadata(mpx_guid)
      if err != nil {
         t.Fatal(err)
      }
      name, err := text.Name(meta)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}

func TestOnDemand(t *testing.T) {
   for _, mpx_guid := range mpx_guids {
      meta, err := NewMetadata(mpx_guid)
      if err != nil {
         t.Fatal(err)
      }
      video, err := meta.OnDemand()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", video)
      time.Sleep(time.Second)
   }
}
