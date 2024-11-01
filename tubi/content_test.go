package tubi

import (
   "41.neocities.org/text"
   "fmt"
   "reflect"
   "testing"
   "time"
)

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Namer{},
   Resolution{},
   VideoContent{},
   VideoResource{},
}

func TestContent(t *testing.T) {
   for _, test := range tests {
      content := &VideoContent{}
      err := content.New(test.content_id, nil)
      if err != nil {
         t.Fatal(err)
      }
      if content.Episode() {
         err := content.New(content.SeriesId, nil)
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

var tests = []struct {
   content_id int
   key_id     string
   url        string
}{
   {
      content_id: 100002888,
      key_id:     "/czNsQXzQQKDN2Bl6kEmDQ==",
      url:        "tubitv.com/movies/100002888",
   },
   {
      content_id: 200042567,
      key_id:     "Ndopo1ozQ8iSL75MAfbL6A==",
      url:        "tubitv.com/tv-shows/200042567",
   },
}
