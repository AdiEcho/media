package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

var id_tests = []string{
   "https://youtube.com/shorts/9Vsdft81Q6w",
   "https://youtube.com/watch?v=XY-hOqcPGCY",
}

const image_test = "UpNXI3_ctAc"

var embed_ids = []string{
   "HtVdAasjOgU",
   "WaOKSUlf4TM",
}

var android_ids = []string{
   "H1BuwMTrtLQ", // content check
   "zv9NimPx3Es",
}

func Test_Android_Embed(t *testing.T) {
   for _, embed_id := range embed_ids {
      var play Player
      var req Request
      req.Android_Embed(embed_id)
      err := play.Post(req, nil)
      if err != nil {
         t.Fatal(err)
      }
      if play.Playability.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}

func Test_Android(t *testing.T) {
   for _, android_id := range android_ids {
      var p Player
      var r Request
      r.Android(android_id)
      err := p.Post(r, nil)
      if err != nil {
         t.Fatal(err)
      }
      if p.Playability.Status != "OK" {
         t.Fatal(p)
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

func Test_ID(t *testing.T) {
   for _, test := range id_tests {
      var req Request
      err := req.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(req.Video_ID)
   }
}

func Test_Image(t *testing.T) {
   for _, img := range Images {
      img.Video_ID = image_test
      fmt.Println(img)
      res, err := http.Head(img.String())
      if err != nil {
         t.Fatal(err)
      }
      if res.StatusCode != http.StatusOK {
         t.Fatal(res.Status)
      }
      time.Sleep(99 * time.Millisecond)
   }
}
