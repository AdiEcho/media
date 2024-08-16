package pluto

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

const default_kid = "0000000063c99438d2d611a908ea7039"

var video_tests = []struct {
   clips string
   start string
   url   string
}{
   {
      url:   "pluto.tv/on-demand/movies/60d9fd1c89632c0013eb2155",
      start: "60d9fd1c89632c0013eb2155",
      clips: "60d9fd1c89632c0013eb2155",
   },
   {
      url:   "pluto.tv/on-demand/movies/la-confidential-1997-1-1",
      start: "la-confidential-1997-1-1",
   },
   {
      url:   "pluto.tv/on-demand/series/65ce5c60a3a8580013c4b64a/episode/65ce5c7ca3a8580013c4be02",
      start: "65ce5c60a3a8580013c4b64a",
      clips: "65ce5c7ca3a8580013c4be02",
   },
   {
      url:   "pluto.tv/on-demand/series/king-of-queens/episode/head-first-1998-1-2",
      start: "king-of-queens",
   },
}

func TestAddress(t *testing.T) {
   for _, test := range video_tests {
      var web Address
      web.Set(test.url)
      fmt.Println(web)
   }
}

///

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
   var pssh widevine.Pssh
   pssh.KeyId, err = hex.DecodeString(default_kid)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Encode())
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(Poster{}, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestClip(t *testing.T) {
   for _, test := range video_tests {
      if test.clips != "" {
         clip, err := Video{Id: test.clips}.Clip()
         if err != nil {
            t.Fatal(err)
         }
         manifest, ok := clip.Dash()
         if !ok {
            t.Fatal("EpisodeClip.Dash")
         }
         url, err := manifest.Parse(Base[0])
         if err != nil {
            t.Fatal(err)
         }
         fmt.Println(url)
         time.Sleep(time.Second)
      }
   }
}

