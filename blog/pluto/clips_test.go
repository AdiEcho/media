package pluto

import (
   "os"
   "testing"
)

func TestClips(t *testing.T) {
   var boot boot_start
   err := boot.New("ex-machina-2015-1-1-ptv1", "99.224.0.0")
   if err != nil {
      t.Fatal(err)
   }
   video, ok := boot.video()
   if !ok {
      t.Fatal("boot_start.video")
   }
   res, err := video.clips()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
