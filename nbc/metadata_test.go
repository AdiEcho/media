package nbc

import (
   "154.pages.dev/encoding"
   "fmt"
   "testing"
   "time"
)

var mpx_guids = []int{
   // episode unlocked
   // nbc.com/saturday-night-live/video/october-21-bad-bunny/9000283422
   9000283422,
   // movie locked
   // nbc.com/cowboys-aliens/video/cowboys-aliens/4340781
   4340781,
}

func TestMetadata(t *testing.T) {
   for _, mpx_guid := range mpx_guids {
      var meta Metadata
      err := meta.New(mpx_guid)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(encoding.Name(meta))
      time.Sleep(time.Second)
   }
}
