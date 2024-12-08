package kanopy

import (
   "fmt"
   "os"
   "testing"
   "time"
)

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
