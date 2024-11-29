package kanopy

import (
   "os"
   "testing"
)

var test = struct{
   key_id string
   url string
   video_id int
}{
   key_id: "DUCS1DH4TB6Po1oEkG9xUA==",
   url: "kanopy.com/product/13808102",
   video_id: 13808102,
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
   resp, err := token.videos(test.video_id)
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
