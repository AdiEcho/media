package kanopy

import (
   "os"
   "testing"
)

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
   for _, test := range tests[:1] {
      resp, err := token.items(test.video_id)
      if err != nil {
         t.Fatal(err)
      }
      defer resp.Body.Close()
      resp.Write(os.Stdout)
   }
}

var tests = []struct{
   key_id string
   url string
   video_id int64
}{
   {
      url: "kanopy.com/irving/video/14881163/14881167",
      video_id: 14881163,
   },
   {
      key_id: "DUCS1DH4TB6Po1oEkG9xUA==",
      url: "kanopy.com/product/13808102",
      video_id: 13808102,
   },
}
