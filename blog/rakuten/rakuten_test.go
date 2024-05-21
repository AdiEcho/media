package rakuten

import (
   "os"
   "testing"
)

func TestFr(t *testing.T) {
   // {"se", "i-heart-huckabees"},
   res, err := streamings(classification["fr"], "jerry-maguire")
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
   //if strings.Contains(text, `"video_quality":"FHD"`) {
}
