package tubi

import (
   "encoding/json"
   "os"
   "testing"
   "time"
)

func TestContent(t *testing.T) {
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   for _, test := range tests {
      var cms content
      err := cms.New(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      if cms.episode() {
         err := cms.New(cms.Series_ID)
         if err != nil {
            t.Fatal(err)
         }
         time.Sleep(time.Second)
      }
      enc.Encode(cms)
   }
}

var tests = map[string]struct{
   content_id int
   key_id string
   url string
}{
   "episode": {
      content_id: 200042567,
      url: "tubitv.com/tv-shows/200042567",
   },
   "movie": {
      content_id: 589292,
      key_id: "943974887f2a4b87a3ded9e99f03c962",
      url: "tubitv.com/movies/589292",
   },
}
