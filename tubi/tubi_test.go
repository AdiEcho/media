package tubi

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestResolution(t *testing.T) {
   for _, test := range tests {
      content := &VideoContent{}
      data, err := content.Marshal(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      err = content.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      if content.Episode() {
         data, err = content.Marshal(content.SeriesId)
         if err != nil {
            t.Fatal(err)
         }
         err = content.Unmarshal(data)
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
      data, err := content.Marshal(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      err = content.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      if content.Episode() {
         data, err = content.Marshal(content.SeriesId)
         if err != nil {
            t.Fatal(err)
         }
         err = content.Unmarshal(data)
         if err != nil {
            t.Fatal(err)
         }
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
