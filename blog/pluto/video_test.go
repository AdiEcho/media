package pluto

import (
   "fmt"
   "testing"
   "time"
)

// seriesIDs:           65ce5c60a3a8580013c4b64a
var video_tests = []struct{
   forward string
   id string
   slug string
   url string
}{
   {
      url: "pluto.tv/on-demand/movies/63c8215d8ba71d0013f29b43",
      id: "63c8215d8ba71d0013f29b43",
      forward: "Canada",
   },
   {
      url: "pluto.tv/on-demand/series/65ce5c60a3a8580013c4b64a/episode/65ce5c7ca3a8580013c4be02",
      id: "65ce5c7ca3a8580013c4be02",
   },
   {
      url: "pluto.tv/on-demand/series/king-of-queens/episode/head-first-1998-1-2",
      slug: "head-first-1998-1-2",
   },
   {
      url: "pluto.tv/on-demand/series/king-of-queens/episode/head-first-1998-1-2",
      slug: "king-of-queens",
   },
   {
      url: "pluto.tv/on-demand/movies/ex-machina-2015-1-1-ptv1",
      slug: "ex-machina-2015-1-1-ptv1",
      forward: "Canada",
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
