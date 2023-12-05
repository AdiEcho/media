package youtube

import (
   "testing"
   "time"
)

func Test_Player_Android(t *testing.T) {
   var req Request
   req.Android()
   req.Video_ID = android_IDs[0]
   for range [9]struct{}{} {
      p, err := req.Player(nil)
      if err != nil {
         t.Fatal(err)
      }
      if len(p.Streaming_Data.Adaptive_Formats) == 0 {
         t.Fatal("adaptiveFormats")
      }
      if p.Video_Details.View_Count == 0 {
         t.Fatal("viewCount")
      }
      time.Sleep(time.Second)
   }
}
