package pluto

import (
   "fmt"
   "testing"
   "time"
)

var video_tests = []struct{
   forward string
   slug string
   url string
}{
   // "pluto.tv/on-demand/series/65ce5c60a3a8580013c4b64a/episode/65ce5c7ca3a8580013c4be02",
   // "pluto.tv/on-demand/series/king-of-queens/episode/head-first-1998-1-2",
   {
      forward: "Canada",
      slug: "63c8215d8ba71d0013f29b43",
      url: "pluto.tv/on-demand/movies/63c8215d8ba71d0013f29b43",
   },
   {
      forward: "Canada",
      slug: "ex-machina-2015-1-1-ptv1",
      url: "pluto.tv/on-demand/movies/ex-machina-2015-1-1-ptv1",
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
