package pluto

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "net/url"
   "os"
   "testing"
   "time"
)

var video_tests = []struct{
   id string
   key_id string
   url   string
}{
   {
      url: "pluto.tv/on-demand/movies/bound-paramount-1-1"
   },
}

func TestAddress(t *testing.T) {
   for _, test := range video_tests {
      var web Address
      web.Set(test.url)
      fmt.Println(web)
   }
}

func TestClip(t *testing.T) {
   base_url, err := url.Parse(Base[len(Base)-1])
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range video_tests {
      clip, err := Video{Id: test.clips}.Clip()
      if err != nil {
         t.Fatal(err)
      }
      manifest, ok := clip.Dash()
      if !ok {
         t.Fatal("EpisodeClip.Dash")
      }
      base_url.Path = manifest.Path
      fmt.Println(base_url)
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

func TestVideo(t *testing.T) {
   for _, test := range video_tests {
      var web Address
      web.Set(test.url)
      video, err := web.Video("")
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", video)
      name, err := text.Name(Namer{video})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
