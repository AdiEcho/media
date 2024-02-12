package nbc

import (
   "154.pages.dev/rosso"
   "fmt"
   "testing"
   "time"
)

var mpx_guids = []int64 {
   // episode unlocked
   // nbc.com/saturday-night-live/video/october-21-bad-bunny/9000283422
   9000283422,
   // movie locked
   // nbc.com/cowboys-aliens/video/cowboys-aliens/4340781
   4340781,
}

func Test_Metadata(t *testing.T) {
   for _, mpx_guid := range mpx_guids {
      meta, err := New_Metadata(mpx_guid)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(rosso.Name(meta))
      time.Sleep(time.Second)
   }
}
