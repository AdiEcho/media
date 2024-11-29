package kanopy

import (
   "net/http"
   "net/url"
   "os"
   "testing"
)

func TestItems(t *testing.T) {
   resp, err := items(14881163)
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
