package paramount

import (
   "fmt"
   "testing"
   "time"
)

func TestItemUsa(t *testing.T) {
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
         fmt.Printf("%+v\n", item)
         time.Sleep(time.Second)
      }
   }
}

func TestItemIntl(t *testing.T) {
   var app AppToken
   err := app.ComCbsCa()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      item, err := app.Item(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      err = item.Unmarshal()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", item)
      time.Sleep(time.Second)
   }
}


