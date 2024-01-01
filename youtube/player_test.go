package youtube

import "testing"

const web_id = "HPkDFc8hq5c"

func Test_Web(t *testing.T) {
   var r Request
   r.Web(web_id)
   var p Player
   err := p.Post(r, nil)
   if err != nil {
      t.Fatal(err)
   }
   if p.Author() == "" {
      t.Fatal("author")
   }
   if p.Playability.Reason != "" {
      t.Fatal("reason")
   }
   if p.Playability.Status != "OK" {
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
   if p.Video_Details.Length_Seconds <= 0 {
      t.Fatal("duration")
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
