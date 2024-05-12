package member

import (
   "os"
   "testing"
)

func TestArticle(t *testing.T) {
   res, err := article()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
