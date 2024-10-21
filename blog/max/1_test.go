package max

import (
   "os"
   "testing"
)

func TestOne(t *testing.T) {
   resp, err := get_one()
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
