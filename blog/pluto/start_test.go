package pluto

import (
   "fmt"
   "testing"
)

func TestStart(t *testing.T) {
   var boot boot_start
   err := boot.New("ex-machina-2015-1-1-ptv1", "99.224.0.0")
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", boot)
}
