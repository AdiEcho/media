package paramount

import (
   "fmt"
   "testing"
   "time"
)

func TestMpdUs(t *testing.T) {
   for _, test := range tests {
      if test.location == "" {
         address, err := Mpd(test.content_id, "DASH_CENC")
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", address)
         time.Sleep(time.Second)
      }
   }
}

func TestMpdFr(t *testing.T) {
   for _, test := range tests {
      if test.location == "France" {
         address, err := Mpd(test.content_id, "DASH_CENC_PRECON")
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", address)
         time.Sleep(time.Second)
      }
   }
}
