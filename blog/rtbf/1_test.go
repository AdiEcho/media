package rtbf

import (
   "fmt"
   "testing"
)

// auvio.rtbf.be/media/i-care-a-lot-i-care-a-lot-3201987
const i_care_a_lot = 3201987

func TestOne(t *testing.T) {
   var embed embed_media
   err := embed.New(i_care_a_lot)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", embed)
}
