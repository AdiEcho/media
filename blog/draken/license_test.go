package draken

import (
   "bytes"
   "fmt"
   "testing"
)

func TestLicense(t *testing.T) {
   res, err := license()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   buf := new(bytes.Buffer)
   res.Write(buf)
   fmt.Printf("%q\n", buf)
}
