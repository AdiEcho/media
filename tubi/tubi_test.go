package tubi

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
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
   for _, test := range tests {
      var pssh widevine.Pssh
      pssh.KeyId, err = base64.StdEncoding.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      content := &VideoContent{}
      content.New(test.content_id)
      content.Unmarshal()
      if content.Episode() {
         content.New(content.SeriesId)
         content.Unmarshal()
         var ok bool
         content, ok = content.Get(test.content_id)
         if !ok {
            t.Fatal("VideoContent.Get")
         }
      }
      video, ok := content.Video()
      if !ok {
         t.Fatal("VideoContent.Video")
      }
      key, err := module.Key(video, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}

func TestResolution(t *testing.T) {
   for _, test := range tests {
      content := &VideoContent{}
      err := content.New(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      err = content.Unmarshal()
      if err != nil {
         t.Fatal(err)
      }
      if content.Episode() {
         err := content.New(content.SeriesId)
         if err != nil {
            t.Fatal(err)
         }
         err = content.Unmarshal()
         if err != nil {
            t.Fatal(err)
         }
         var ok bool
         content, ok = content.Get(test.content_id)
         if !ok {
            t.Fatal("VideoContent.Get")
         }
      }
      fmt.Println(test.url)
      for _, v := range content.VideoResources {
         fmt.Println(v.Resolution, v.Type)
      }
      time.Sleep(time.Second)
   }
}
