package paramount

import (
   "fmt"
   "testing"
   "time"
)

func TestMpdUs(t *testing.T) {
   var app AppToken
   err := app.ComCbsApp()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      if test.location == "" {
         item, err := app.Item(test.content_id)
         if err != nil {
            t.Fatal(err)
         }
         err = item.Unmarshal()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", item.Mpd())
         time.Sleep(time.Second)
      }
   }
}

func TestMpdFr(t *testing.T) {
   var app AppToken
   err := app.ComCbsCa()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      if test.location == "France" {
         item, err := app.Item(test.content_id)
         if err != nil {
            t.Fatal(err)
         }
         err = item.Unmarshal()
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", item.Mpd())
         time.Sleep(time.Second)
      }
   }
}
