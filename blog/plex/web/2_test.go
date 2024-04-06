package plex

import (
   "os"
   "testing"
)

func TestMetadata(t *testing.T) {
   var anon anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   res, err := anon.metadata("https://watch.plex.tv/movie/cruel-intentions")
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
