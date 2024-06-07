package rtbf

import (
   "fmt"
   "testing"
   "time"
)

var media = map[string]struct{
   id int64
   key_id string
   url string
}{
   "film": {
      id: 3201987,
      key_id: "o1C37Tt5SzmHMmEgQViUEA==",
      url: "auvio.rtbf.be/media/i-care-a-lot-i-care-a-lot-3201987",
   },
   "episode": {
      id: 3194636,
      url: "auvio.rtbf.be/media/grantchester-grantchester-s01-3194636",
   },
}

func TestEmbedMedia(t *testing.T) {
   for _, medium := range media {
      var embed embed_media
      err := embed.New(medium.id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", embed)
      time.Sleep(time.Second)
   }
}
