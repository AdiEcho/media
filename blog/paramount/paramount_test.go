package paramount

import (
   "fmt"
   "testing"
)

func TestItem(t *testing.T) {
   var app AppToken
   err := app.New()
   if err != nil {
      t.Fatal(err)
   }
   item, err := app.Item(content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
}

// paramountplus.com/movies/video/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ
const content_id = "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ"

func TestMpd(t *testing.T) {
   address, err := DashCenc(content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(address)
}
