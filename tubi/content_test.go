package tubi

import (
   "41.neocities.org/text"
   "fmt"
   "reflect"
   "testing"
   "time"
)

func TestContent(t *testing.T) {
   for _, test := range tests {
      content := &VideoContent{}
      data, err := content.Marshal(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      err = content.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      if content.Episode() {
         data, err = content.Marshal(content.SeriesId)
         if err != nil {
            t.Fatal(err)
         }
         err = content.Unmarshal(data)
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
