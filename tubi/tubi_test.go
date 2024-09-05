package tubi

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestLicense(t *testing.T) {
   test := tests["the-mask"]
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
   pssh.KeyId, err = hex.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      t.Fatal(err)
   }
   var content VideoContent
   err = content.New(test.content_id)
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(content.Video(), pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestResolution(t *testing.T) {
   for _, test := range tests {
      content := &VideoContent{}
      err := content.New(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      if content.Episode() {
         err := content.New(content.SeriesId)
         if err != nil {
            t.Fatal(err)
         }
         time.Sleep(time.Second)
         var ok bool
         content, ok = content.Get(test.content_id)
         if !ok {
            t.Fatal("get")
         }
      }
      fmt.Println(test.url)
      for _, r := range content.VideoResources {
         fmt.Println(r.Resolution, r.Type)
      }
   }
}
