package pluto

import (
   "154.pages.dev/text"
   "41.neocities.org/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "reflect"
   "testing"
   "time"
)

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Address{},
   EpisodeClip{},
   FileBase{},
   Namer{},
   OnDemand{},
   Poster{},
   Url{},
   VideoSeason{},
}

func TestAddress(t *testing.T) {
   for _, test := range video_tests {
      var web Address
      err := web.Set(test.url)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(web)
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
      manifest.Scheme = Base[0].Scheme
      manifest.Host = Base[0].Host
      fmt.Printf("%+v\n", manifest)
      time.Sleep(time.Second)
   }
}

// the slug is useful as it sometimes contains the year, but its not worth
// parsing since its sometimes missing
var video_tests = []struct {
   id     string
   key_id string
   url    string
}{
   {
      id:     "5c4bb2b308d10f9a25bbc6af",
      key_id: "0000000066bfe3cd26602c92dc082e3b",
      url:    "pluto.tv/on-demand/movies/bound-paramount-1-1",
   },
   {
      id:     "66b3838317101c00130b411e",
      key_id: "0000000066b3c161c1cee84ffce71de3",
      url:    "pluto.tv/on-demand/movies/just-go-with-it-2011-1-1",
   },
   {
      id:     "6356d14136d64a001450b121",
      key_id: "000000006358c035248b647dad3c09ad",
      url:    "pluto.tv/on-demand/series/frasier-cbs-tv/season/1/episode/space-quest-1992-1-2",
   },
}
