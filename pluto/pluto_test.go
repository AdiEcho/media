package pluto

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "reflect"
   "testing"
   "time"
)

// the slug is useful as it sometimes contains the year, but its not worth
// parsing since its sometimes missing
var video_tests = []struct{
   id     string
   key_id string
   url    string
}{
   {
      id:     "675a0fa22678a50014690c3f",
      key_id: "AAAAAGdaD6FuwTSRB/+yHg==",
      url:    "pluto.tv/on-demand/movies/675a0fa22678a50014690c3f",
   },
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
   Client{},
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
      name := text.Name(Namer{video})
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
      var module widevine.Module
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(Client{}, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}
