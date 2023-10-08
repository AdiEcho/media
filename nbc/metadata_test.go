package nbc

import (
   "154.pages.dev/media"
   "fmt"
   "testing"
   "time"
)

var guids = []int64{
   // episode
   // nbc.com/the-irrational/video/dead-woman-walking/9000360354
   9000360354,
   // episode
   // nbc.com/pasion-de-gavilanes/video/una-verguenza/3760495
   3760495,
   // movie
   // nbc.com/john-wick/video/john-wick/3448375
   3448375,
}

func Test_Meta(t *testing.T) {
   for _, guid := range guids {
      meta, err := New_Metadata(guid)
      if err != nil {
         t.Fatal(err)
      }
      name, err := media.Name(meta)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
}
