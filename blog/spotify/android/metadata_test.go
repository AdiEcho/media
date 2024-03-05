package android

import (
   "bytes"
   "fmt"
   "os"
   "testing"
)

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
