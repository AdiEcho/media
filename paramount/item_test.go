package paramount

import (
   "41.neocities.org/text"
   "fmt"
   "testing"
   "time"
)

func TestMpdUs(t *testing.T) {
   var token AppToken
   err := token.ComCbsApp()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      if test.location == "" {
         var item VideoItem
         data, err := item.Marshal(token, test.content_id)
         if err != nil {
            t.Fatal(err)
         }
         err = item.Unmarshal(data)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", item.Mpd())
         time.Sleep(time.Second)
      }
   }
}

func TestMpdFr(t *testing.T) {
   var token AppToken
   err := token.ComCbsCa()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      if test.location == "France" {
         var item VideoItem
         data, err := item.Marshal(token, test.content_id)
         if err != nil {
            t.Fatal(err)
         }
         err = item.Unmarshal(data)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", item.Mpd())
         time.Sleep(time.Second)
      }
   }
}

func TestItemUsa(t *testing.T) {
   var token AppToken
   err := token.ComCbsApp()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      if test.location == "" {
         var item VideoItem
         data, err := item.Marshal(token, test.content_id)
         if err != nil {
            t.Fatal(err)
         }
         err = item.Unmarshal(data)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", text.Name(&item))
         time.Sleep(time.Second)
      }
   }
}
