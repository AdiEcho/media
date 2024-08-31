package paramount

import (
   "fmt"
   "testing"
)

func TestItemUs(t *testing.T) {
   var app AppToken
   err := app.ComCbsApp()
   if err != nil {
      t.Fatal(err)
   }
   item, err := app.Item(tests["us"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   err = item.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
}

func TestItemFr(t *testing.T) {
   var app AppToken
   err := app.ComCbsCa()
   if err != nil {
      t.Fatal(err)
   }
   item, err := app.Item(tests["fr"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   err = item.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
}
