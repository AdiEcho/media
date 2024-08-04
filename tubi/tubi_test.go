package tubi

import (
   "154.pages.dev/text"
   "fmt"
   "testing"
   "time"
)

func TestContent(t *testing.T) {
   for _, test := range tests {
      cms := &Content{}
      err := cms.New(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(time.Second)
      if cms.Episode() {
         err := cms.New(cms.SeriesId)
         if err != nil {
            t.Fatal(err)
         }
         time.Sleep(time.Second)
         var ok bool
         cms, ok = cms.Get(test.content_id)
         if !ok {
            t.Fatal("get")
         }
      }
      name, err := text.Name(Namer{cms})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
   }
}

var tests = map[string]struct {
   content_id int
   key_id     string
   url        string
}{
   "the-mask": {
      content_id: 589292,
      key_id:     "943974887f2a4b87a3ded9e99f03c962",
      url:        "tubitv.com/movies/589292",
   },
   "the-devil-s-advocate": {
      content_id: 590133,
      url:        "tubitv.com/movies/590133",
   },
   "training-day": {
      content_id: 608618,
      url:        "tubitv.com/movies/608618",
   },
   "heat": {
      content_id: 611324,
      url:        "tubitv.com/movies/611324",
   },
   "the-thing": {
      content_id: 648069,
      url:        "tubitv.com/movies/648069",
   },
   "get-out": {
      content_id: 100009092,
      url:        "tubitv.com/movies/100009092",
   },
   "lady-bird": {
      content_id: 100012358,
      url:        "tubitv.com/movies/100012358",
   },
   "the-batman": {
      content_id: 100016185,
      url:        "tubitv.com/movies/100016185",
   },
   "scandal": {
      content_id: 200042567,
      url:        "tubitv.com/tv-shows/200042567",
   },
}
