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
      url: "kanopy.com/en/product/13808102",
      video_id: 13808102,
   },
   {
      url: "kanopy.com/en/product/14881161",
      video_id: 14881161,
   },
   {
      url: "kanopy.com/en/product/14881167",
      video_id: 14881167,
   },
}

func TestVideos(t *testing.T) {
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
      video, err := token.videos(test.video_id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", video)
      time.Sleep(time.Second)
   }
}
