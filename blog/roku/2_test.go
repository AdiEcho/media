package roku

import (
   "os"
   "testing"
)

const roku_id = "597a64a4a25c5bf6af4a8c7053049a6f"

func TestPlayback(t *testing.T) {
   var token account_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   res, err := token.playback(roku_id)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
