package youtube

import "testing"

const web_id = "HPkDFc8hq5c"

func TestWeb(t *testing.T) {
   r := Request{VideoId: web_id}
   r.Web()
   var p Player
   err := p.Post(r, nil)
   if err != nil {
      t.Fatal(err)
   }
   if p.Author() == "" {
      t.Fatal("author")
   }
   if p.PlayabilityStatus.Reason != "" {
      t.Fatal("reason")
   }
   if p.PlayabilityStatus.Status != "OK" {
      t.Fatal("status")
   }
   if len(p.StreamingData.AdaptiveFormats) == 0 {
      t.Fatal("adaptiveFormats")
   }
   if _, err := p.Time(); err != nil {
      t.Fatal(err)
   }
   if p.Title() == "" {
      t.Fatal("title")
   }
   if p.VideoDetails.LengthSeconds <= 0 {
      t.Fatal("duration")
   }
   if p.VideoDetails.ShortDescription == "" {
      t.Fatal("shortDescription")
   }
   if p.VideoDetails.VideoId == "" {
      t.Fatal("videoId")
   }
   if p.VideoDetails.ViewCount <= 0 {
      t.Fatal("viewCount")
   }
}
