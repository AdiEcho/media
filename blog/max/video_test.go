package max

import (
   "os"
   "testing"
)

func TestVideo(t *testing.T) {
   var st st_cookie
   err := st.New()
   if err != nil {
      t.Fatal(err)
   }
   resp, err := st.video()
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
