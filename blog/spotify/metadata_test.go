package spotify

import (
   "bytes"
   "fmt"
   "testing"
)

func TestMetadata(t *testing.T) {
   var ok login_ok
   res, err := ok.metadata()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   var b bytes.Buffer
   res.Write(&b)
   fmt.Printf("%q\n", b.Bytes())
}
