package nbc

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
   CoreVideo{},
   Metadata{},
   OnDemand{},
   page_request{},
}

func TestMetadata(t *testing.T) {
   for _, test := range tests {
      var meta Metadata
      err := meta.New(test.id)
      if err != nil {
         t.Fatal(err)
      }
      name, err := text.Name(&meta)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}

var tests = []struct{
   url string
   program string
   id int
   lock bool
   key_id string
}{
   {
      id: 9000283422,
      key_id: "0552e44842654a4e81b326004be47be0",
      program: "episode",
      url: "nbc.com/saturday-night-live/video/may-18-jake-gyllenhaal/9000283438",
   },
   {
      id: 9000283435,
      key_id: "a48d84f23ec74aa1ba8b1d4c863ac02b",
      lock: true,
      program: "episode",
      url: "nbc.com/saturday-night-live/video/march-30-ramy-youssef/9000283435",
   },
   {
      id: 2957739,
      key_id: "e416811be8c44b8e9e598ea7b22e57cc",
      lock: true,
      program: "movie",
      url: "nbc.com/2-fast-2-furious/video/2-fast-2-furious/2957739",
   },
}
