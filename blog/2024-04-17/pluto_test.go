package pluto

import (
   "fmt"
   "testing"
)

func TestBoot(t *testing.T) {
   var boot boot_start
   err := boot.New("ex-machina-2015-1-1-ptv1", forwards["Canada"])
   if err != nil {
      t.Fatal(err)
   }
   mpd, ok := boot.mpd()
   fmt.Println(mpd, ok)
}
