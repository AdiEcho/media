package nbc

import (
   "154.pages.dev/text"
   "fmt"
   "testing"
   "time"
)

var mpx_guids = []int{
   // episode unlocked
   // nbc.com/saturday-night-live/video/may-18-jake-gyllenhaal/9000283438
   9000283422,
   // movie locked
   // nbc.com/2-fast-2-furious/video/2-fast-2-furious/2957739
   2957739,
}

func TestMetadata(t *testing.T) {
   for _, mpx_guid := range mpx_guids {
      var meta Metadata
      err := meta.New(mpx_guid)
      if err != nil {
         t.Fatal(err)
      }
      name, err := text.Name(meta)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
