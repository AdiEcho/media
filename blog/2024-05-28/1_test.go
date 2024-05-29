package roku

import (
   "os"
   "testing"
)

func TestOne(t *testing.T) {
   res, err := one()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
