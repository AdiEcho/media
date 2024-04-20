package pluto

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

var video_tests = []struct{
   forward string
   slug string
   url string
}{
   {
      url: "pluto.tv/on-demand/movies/63c8215d8ba71d0013f29b43",
      slug: "63c8215d8ba71d0013f29b43",
      forward: "Canada",
   },
   {
      url: "pluto.tv/on-demand/movies/ex-machina-2015-1-1-ptv1",
      slug: "ex-machina-2015-1-1-ptv1",
      forward: "Canada",
   },
   {
      url: "pluto.tv/on-demand/series/65ce5c60a3a8580013c4b64a/episode/65ce5c7ca3a8580013c4be02",
      slug: "65ce5c7ca3a8580013c4be02",
   },
   {
      url: "pluto.tv/on-demand/series/king-of-queens/episode/head-first-1998-1-2",
      slug: "head-first-1998-1-2",
   },
}

func TestVideo(t *testing.T) {
   for _, test := range video_tests {
      video, err := new_video(test.slug, forwards[test.forward])
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", video)
      time.Sleep(time.Second)
   }
}
var addresses = []string{
   // pluto.tv/on-demand/series/king-of-queens/episode/head-first-1998-1-2
   // pluto.tv/on-demand/series/65ce5c60a3a8580013c4b64a/episode/65ce5c7ca3a8580013c4be02
   "65ce5c7ca3a8580013c4be02",
   // pluto.tv/on-demand/movies/ex-machina-2015-1-1-ptv1
   // pluto.tv/on-demand/movies/63c8215d8ba71d0013f29b43
   "63c8215d8ba71d0013f29b43",
}

func TestClip(t *testing.T) {
   for _, address := range addresses {
      clip, err := new_clip(address)
      if err != nil {
         t.Fatal(err)
      }
      manifest, ok := clip.dash()
      if !ok {
         t.Fatal("episode_clip.dash")
      }
      url, err := manifest.parse(bases[0])
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(url)
      time.Sleep(time.Second)
   }
}

const default_kid = "0000000063c99438d2d611a908ea7039"

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
   key_id, err := hex.DecodeString(default_kid)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   if err := module.New(private_key, client_id, key_id); err != nil {
      t.Fatal(err)
   }
   license, err := module.License(poster{})
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Printf("%x %v\n", key, ok)
}
