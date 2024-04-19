package pluto

import (
   "fmt"
   "testing"
)

func TestDash(t *testing.T) {
   video, err := new_video("ex-machina-2015-1-1-ptv1", "99.224.0.0")
   if err != nil {
      t.Fatal(err)
   }
   clip, err := video.clip()
   if err != nil {
      t.Fatal(err)
   }
   manifest, ok := clip.dash()
   if !ok {
      t.Fatal("episode_clip.dash")
   }
   url, err := manifest.parse()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(url)
}

func TestHls(t *testing.T) {
   video, err := new_video("ex-machina-2015-1-1-ptv1", "99.224.0.0")
   if err != nil {
      t.Fatal(err)
   }
   clip, err := video.clip()
   if err != nil {
      t.Fatal(err)
   }
   manifest, ok := clip.hls()
   if !ok {
      t.Fatal("episode_clip.hls")
   }
   url, err := manifest.parse()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(url)
}
