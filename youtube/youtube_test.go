package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

func TestAndroid(t *testing.T) {
   for _, android_id := range android_ids {
      var p Player
      r := Request{VideoId: android_id}
      r.Android()
      err := p.Post(r, nil)
      if err != nil {
         t.Fatal(err)
      }
      if p.PlayabilityStatus.Status != "OK" {
         t.Fatal(p)
      }
      if len(p.StreamingData.AdaptiveFormats) == 0 {
         t.Fatal("adaptiveFormats")
      }
      if p.VideoDetails.ViewCount == 0 {
         t.Fatal("viewCount")
      }
      time.Sleep(time.Second)
   }
}

func TestId(t *testing.T) {
   for _, test := range id_tests {
      var req Request
      err := req.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(req.VideoId)
   }
}

func TestImage(t *testing.T) {
   for _, img := range Images {
      img.VideoId = image_test
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

func TestAndroidEmbed(t *testing.T) {
   for _, embed_id := range embed_ids {
      var play Player
      req := Request{VideoId: embed_id}
      req.AndroidEmbed()
      err := play.Post(req, nil)
      if err != nil {
         t.Fatal(err)
      }
      if play.PlayabilityStatus.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}

