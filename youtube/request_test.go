package youtube

import (
   "154.pages.dev/http"
   "os"
   "testing"
   "time"
)

const web_ID = "HPkDFc8hq5c"

func Test_Web(t *testing.T) {
   http.No_Location()
   http.Verbose()
   req := Request{Video_ID: web_ID}
   req.Web()
   p, err := req.Player(nil)
   if err != nil {
      t.Fatal(err)
   }
   if p.Author() == "" {
      t.Fatal("author")
   }
   if p.Duration() <= 0 {
      t.Fatal("duration")
   }
   if p.Playability_Status.Reason != "" {
      t.Fatal("reason")
   }
   if p.Playability_Status.Status != "OK" {
      t.Fatal("status")
   }
   if len(p.Streaming_Data.Adaptive_Formats) == 0 {
      t.Fatal("adaptiveFormats")
   }
   if _, err := p.Time(); err != nil {
      t.Fatal(err)
   }
   if p.Title() == "" {
      t.Fatal("title")
   }
   if p.Video_Details.Short_Description == "" {
      t.Fatal("shortDescription")
   }
   if p.Video_Details.Video_ID == "" {
      t.Fatal("videoId")
   }
   if p.Video_Details.View_Count <= 0 {
      t.Fatal("viewCount")
   }
}

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
      if play.Playability_Status.Status != "OK" {
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

func Test_Android_Check(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   tok, err := Read_Token(home + "/youtube.json")
   if err != nil {
      t.Fatal(err)
   }
   var req Request
   req.Android_Check()
   for _, check := range check_IDs {
      req.Video_ID = check
      play, err := req.Player(tok)
      if err != nil {
         t.Fatal(err)
      }
      if play.Playability_Status.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
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
      play, err := req.Player(nil)
      if err != nil {
         t.Fatal(err)
      }
      if play.Playability_Status.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}
