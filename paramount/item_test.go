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
   items, err := app.Items(tests["us"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", items)
}

func TestItemFr(t *testing.T) {
   var app AppToken
   err := app.ComCbsCa()
   if err != nil {
      t.Fatal(err)
   }
   items, err := app.Items(tests["fr"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", items)
}
