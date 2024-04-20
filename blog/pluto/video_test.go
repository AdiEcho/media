package pluto

import (
   "fmt"
   "testing"
   "time"
)

// seriesIDs:           65ce5c60a3a8580013c4b64a
var video_tests = []struct{
   id string
   slug string
   url string
}{
   {
      url: "pluto.tv/on-demand/movies/60d9fd1c89632c0013eb2155",
      id: "60d9fd1c89632c0013eb2155",
   },
   {
      url: "pluto.tv/on-demand/movies/la-confidential-1997-1-1",
      slug: "la-confidential-1997-1-1",
   },
   {
      url: "pluto.tv/on-demand/series/65ce5c60a3a8580013c4b64a/episode/65ce5c7ca3a8580013c4be02",
      id: "65ce5c60a3a8580013c4b64a",
   },
   {
      url: "pluto.tv/on-demand/series/king-of-queens/episode/head-first-1998-1-2",
      slug: "king-of-queens",
   },
}

func TestVideo(t *testing.T) {
   for _, test := range video_tests {
      input, ok := func() (string, bool) {
         switch {
         case test.id != "":
            return test.id, true
         case test.slug != "":
            return test.slug, true
         }
         return "", false
      }()
      if !ok {
         t.Fatal(test)
      }
      video, err := new_video(input)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", video)
      time.Sleep(time.Second)
   }
}
