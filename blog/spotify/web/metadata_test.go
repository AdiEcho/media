package web

import (
   "bytes"
   "fmt"
   "os"
   "testing"
)

req.URL.Path = "/metadata/4/track/2da9a11032664413b24de181c534f157"

func TestMetadata(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var login LoginOk
   login.Data, err = os.ReadFile(home + "/spotify.bin")
   if err != nil {
      t.Fatal(err)
   }
   if err := login.Consume(); err != nil {
      t.Fatal(err)
   }
   res, err := login.metadata()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   var b bytes.Buffer
   res.Write(&b)
   fmt.Printf("%q\n", b.Bytes())
}
