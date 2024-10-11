package paramount

import (
   "fmt"
   "testing"
   "time"
)

func TestMpdUsa(t *testing.T) {
   for _, test := range tests {
      if test.location == "" {
         mpd, err := Location(test.content_id, false)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", mpd)
         time.Sleep(time.Second)
      }
   }
}

func TestMpdIntl(t *testing.T) {
   for _, test := range tests {
      if test.location != "" {
         mpd, err := Location(test.content_id, true)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", mpd)
         time.Sleep(time.Second)
      }
   }
}
