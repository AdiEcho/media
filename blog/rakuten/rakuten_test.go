package rakuten

import (
   "os"
   "testing"
)

//if strings.Contains(text, `"video_quality":"FHD"`) {

func TestFr(t *testing.T) {
   res, err := streamings(classification["fr"], "jerry-maguire")
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

func TestSe(t *testing.T) {
   res, err := streamings(classification["se"], "i-heart-huckabees")
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
