package rtbf

import (
   "os"
   "testing"
)

// auvio.rtbf.be/media/i-care-a-lot-i-care-a-lot-3201987
const i_care_a_lot = 3201987

func TestOne(t *testing.T) {
   res, err := one(i_care_a_lot)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
