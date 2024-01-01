package youtube

import (
   "154.pages.dev/stream"
   "fmt"
   "testing"
   "time"
)

// youtube.com/channel/UCuVPpxrm2VAgpH3Ktln4HXg
var ids = []string{
   "7KLCti7tOXE", // video
   "2ZcDwdXEVyI", // episode
   "PBcnZCa1dEk", // film
}

func Test_Next(t *testing.T) {
   for _, id := range ids {
      var con Contents
      req := Request{Video_ID: id}
      req.Web()
      err := con.Next(req)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(stream.Name(con))
      time.Sleep(99*time.Millisecond)
   }
}
