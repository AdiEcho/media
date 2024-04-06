package plex

import (
   "os"
   "testing"
)

const cruel = "https://watch.plex.tv/movie/cruel-intentions"

func TestMetadata(t *testing.T) {
   var anon anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   res, err := anon.metadata(cruel)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
