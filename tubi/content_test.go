package tubi

import (
   "41.neocities.org/text"
   "fmt"
   "testing"
   "time"
)

var tests = []struct{
   content_id int
   key_id     string
   url        string
}{
   {
      url:        "tubitv.com/movies/590133",
      key_id: "8qyB6sGARQWT++zcgNlnwg==",
      content_id: 590133,
   },
   {
      url:        "tubitv.com/tv-shows/200042567",
      key_id: "Ndopo1ozQ8iSL75MAfbL6A==",
      content_id: 200042567,
   },
}

func TestContent(t *testing.T) {
   for _, test := range tests {
      content := &VideoContent{}
      err := content.New(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      err = content.Unmarshal()
      if err != nil {
         t.Fatal(err)
      }
      if content.Episode() {
         err := content.New(content.SeriesId)
         if err != nil {
            t.Fatal(err)
         }
         err = content.Unmarshal()
         if err != nil {
            t.Fatal(err)
         }
         var ok bool
         content, ok = content.Get(test.content_id)
         if !ok {
            t.Fatal("VideoContent.Get")
         }
      }
      name, err := text.Name(Namer{content})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
