package youtube

import (
   "fmt"
   "testing"
   "time"
)

// youtube.com/channel/UCuVPpxrm2VAgpH3Ktln4HXg
var video_ids = []string{
   "7KLCti7tOXE", // video
   "2ZcDwdXEVyI", // episode
   "PBcnZCa1dEk", // film
}

func TestNext(t *testing.T) {
   for _, video_id := range video_ids {
      req := Request{Video_ID: video_id}
      req.Web()
      var con Contents
      err := con.Next(req)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(stream.Name(con))
      time.Sleep(99*time.Millisecond)
   }
}
