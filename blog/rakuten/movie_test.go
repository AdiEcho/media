package rakuten

import (
   "os"
   "testing"
)

func TestMovie(t *testing.T) {
   res, err := gizmo_movie()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
