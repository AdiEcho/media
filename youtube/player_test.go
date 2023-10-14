package youtube

import (
   "154.pages.dev/http"
   "testing"
   "time"
)

func Test_Player_Web(t *testing.T) {
   var req Request
   req.Mobile_Web()
   req.Video_ID = androids[0]
   http.No_Location()
   http.Verbose()
   for range [9]struct{}{} {
      p, err := req.Player(nil)
      if err != nil {
         t.Fatal(err)
      }
      if _, err := p.Time(); err != nil {
         t.Fatal(err)
      }
      if p.Video_Details.View_Count == 0 {
         t.Fatal("viewCount")
      }
      time.Sleep(time.Second)
   }
}

func Test_Player_Android(t *testing.T) {
   var req Request
   req.Android()
   req.Video_ID = androids[0]
   http.No_Location()
   http.Verbose()
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
