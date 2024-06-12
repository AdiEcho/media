package max

import (
   "os"
   "testing"
)

func TestPlayback(t *testing.T) {
   resp, err := playback()
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
