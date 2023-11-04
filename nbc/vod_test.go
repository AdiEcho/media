package nbc

import (
   "fmt"
   "testing"
)

func Test_VOD(t *testing.T) {
   for _, mpx_guid := range mpx_guids {
      fmt.Println(mpx_guid)
   }
}

var mpx_guids = []int64 {
   // locked content
   // nbc.com/saturday-night-live/video/october-28-nate-bargatze/9000283426
   9000283426,
   // unlocked content
   // nbc.com/saturday-night-live/video/october-21-bad-bunny/9000283422
   9000283422,
}
