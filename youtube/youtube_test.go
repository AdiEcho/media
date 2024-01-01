package youtube

import (
   "fmt"
   "net/http"
   "testing"
   "time"
)

const web_ID = "HPkDFc8hq5c"

var embed_IDs = []string{
   "HtVdAasjOgU",
   "WaOKSUlf4TM",
}

func Test_Android_Embed(t *testing.T) {
   var req Request
   req.Android_Embed()
   for _, embed := range embed_IDs {
      req.Video_ID = embed
      play, err := req.Player(nil)
      if err != nil {
         t.Fatal(err)
      }
      if play.Playability.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}

var check_IDs = []string{
   "Cr381pDsSsA", // racy check
   "HsUATh_Nc2U", // racy check
   "SZJvDhaSDnc", // racy check
   "Tq92D6wQ1mg", // racy check
   "dqRZDebPIGs", // racy check
   "nGC3D_FkCmg", // content check
}

var android_IDs = []string{
   "H1BuwMTrtLQ", // content check
   "zv9NimPx3Es",
}

func Test_Android(t *testing.T) {
   var req Request
   req.Android()
   for _, android := range android_IDs {
      req.Video_ID = android
      p, err := req.Player(nil)
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

const image_test = "UpNXI3_ctAc"

func Test_Image(t *testing.T) {
   req, err := http.NewRequest("HEAD", "", nil)
   if err != nil {
      t.Fatal(err)
   }
   for _, img := range Images {
      req.URL = img.URL(image_test)
      fmt.Println("HEAD", req.URL)
      res, err := new(http.Transport).RoundTrip(req)
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

func Test_ID(t *testing.T) {
   var req Request
   for _, test := range id_tests {
      err := req.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(req.Video_ID)
   }
}
