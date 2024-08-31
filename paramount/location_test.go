package paramount

import (
   "fmt"
   "testing"
)

func TestMpdFr(t *testing.T) {
   mpd, err := Location(tests["fr"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", mpd)
}

func TestMpdUs(t *testing.T) {
   mpd, err := Location(tests["us"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", mpd)
}
