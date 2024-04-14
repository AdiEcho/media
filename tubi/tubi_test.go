package tubi

import (
   "154.pages.dev/media/internal"
   "fmt"
   "testing"
   "time"
)

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

func TestContent(t *testing.T) {
   for _, test := range tests {
      cms := new(content)
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
         var ok bool
         cms, ok = cms.get(test.content_id)
         if !ok {
            t.Fatal("get")
         }
      }
      name, err := internal.Name(namer{cms})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
   }
}
