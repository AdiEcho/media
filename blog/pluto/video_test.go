package pluto

import (
   "fmt"
   "testing"
   "time"
)

var video_tests = []struct{
   clips string
   start string
   url string
}{
   {
      clips: "60d9fd1c89632c0013eb2155",
      start: "60d9fd1c89632c0013eb2155",
      url: "pluto.tv/on-demand/movies/60d9fd1c89632c0013eb2155",
   },
   {
      start: "la-confidential-1997-1-1",
      url: "pluto.tv/on-demand/movies/la-confidential-1997-1-1",
   },
   {
      clips: "65ce5c7ca3a8580013c4be02",
      start: "65ce5c60a3a8580013c4b64a",
      url: "pluto.tv/on-demand/series/65ce5c60a3a8580013c4b64a/episode/65ce5c7ca3a8580013c4be02",
   },
   {
      start: "king-of-queens",
      url: "pluto.tv/on-demand/series/king-of-queens/episode/head-first-1998-1-2",
   },
}

func TestVideo(t *testing.T) {
   for _, test := range video_tests {
      video, err := new_video(test.start)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", video)
      time.Sleep(time.Second)
   }
}
