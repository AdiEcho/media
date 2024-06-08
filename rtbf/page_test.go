package rtbf

import (
   "os"
   "testing"
   "time"
)

func TestPage(t *testing.T) {
   for _, medium := range media {
      func() {
         res, err := page(medium.path)
         if err != nil {
            t.Fatal(err)
         }
         defer res.Body.Close()
         res.Write(os.Stdout)
      }()
      time.Sleep(time.Second)
   }
}
