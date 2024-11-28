package kanopy

import (
   "os"
   "testing"
)

func TestPlays(t *testing.T) {
   resp, err := plays(test.video_id)
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
