package plex

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = map[string]Path{
   "episode": {"/show/broadchurch/season/3/episode/5"},
   // watch.plex.tv/movie/cruel-intentions
   "movie": {"/movie/cruel-intentions"},
}

func TestDiscover(t *testing.T) {
   var anon Anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      match, err := anon.Discover(test)
      if err != nil {
         t.Fatal(err)
      }
      name, err := text.Name(Namer{match})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
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
   var pssh widevine.Pssh
   pssh.KeyId, err = hex.DecodeString(default_kid)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      t.Fatal(err)
   }
   var anon Anonymous
   err = anon.New()
   if err != nil {
      t.Fatal(err)
   }
   match, err := anon.Discover(tests["movie"])
   if err != nil {
      t.Fatal(err)
   }
   video, err := anon.Video(match, "")
   if err != nil {
      t.Fatal(err)
   }
   part, ok := video.Dash(anon)
   if !ok {
      t.Fatal("Metadata.Dash")
   }
   key, err := module.Key(part, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

const default_kid = "eabdd790d9279b9699b32110eed9a154"
