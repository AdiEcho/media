package max

import (
   "os"
   "testing"
)

func TestPlayback(t *testing.T) {
   text, err := os.ReadFile("login.json")
   if err != nil {
      t.Fatal(err)
   }
   var st st_cookie
   err = st.unmarshal(text)
   if err != nil {
      t.Fatal(err)
   }
   var playback playback_request
   playback.New()
   resp, err := st.playback(playback)
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
