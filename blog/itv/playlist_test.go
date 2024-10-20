package itv

import (
   "os"
   "testing"
)

func TestPlaylist(t *testing.T) {
   resp, err := playlist()
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
