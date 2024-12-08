package kanopy

import (
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = []struct{
   key_id string
   url string
   video_id int64
}{
   {
      key_id: "DUCS1DH4TB6Po1oEkG9xUA==",
      url: "kanopy.com/irving/video/13808102",
      video_id: 13808102,
   },
   {
      url: "kanopy.com/irving/video/14881163/14881167",
      video_id: 14881163,
   },
}

func TestItems(t *testing.T) {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token web_token
   err = token.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      items, err := token.items(test.video_id)
      if err != nil {
         t.Fatal(err)
      }
      for _, item := range items.List {
         fmt.Printf("%+v\n", item)
      }
      time.Sleep(time.Second)
   }
}
