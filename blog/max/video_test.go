package max

import (
   "os"
   "testing"
)

func TestVideo(t *testing.T) {
   var token default_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   resp, err := token.video()
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
