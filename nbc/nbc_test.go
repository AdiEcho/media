package nbc

import (
   "154.pages.dev/http"
   "154.pages.dev/stream"
   "fmt"
   "testing"
   "time"
)

func Test_On_Demand(t *testing.T) {
   http.No_Location()
   http.Verbose()
   for _, mpx_guid := range mpx_guids {
      meta, err := New_Metadata(mpx_guid)
      if err != nil {
         t.Fatal(err)
      }
      video, err := meta.On_Demand()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", video)
      time.Sleep(time.Second)
   }
}

var mpx_guids = []int64 {
   // episode locked
   // nbc.com/saturday-night-live/video/october-28-nate-bargatze/9000283426
   9000283426,
   // episode unlocked
   // nbc.com/saturday-night-live/video/october-21-bad-bunny/9000283422
   9000283422,
   // movie locked
   // nbc.com/john-wick/video/john-wick/3448375
   3448375,
}

func Test_Meta(t *testing.T) {
   for _, mpx_guid := range mpx_guids {
      meta, err := New_Metadata(mpx_guid)
      if err != nil {
         t.Fatal(err)
      }
      name, err := stream.Format_Film(meta)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
}

