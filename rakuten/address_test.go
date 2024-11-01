package rakuten

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
   Address{},
   GizmoMovie{},
   OnDemand{},
   StreamInfo{},
}

func TestMovie(t *testing.T) {
   for _, test := range tests {
      var web Address
      err := web.Set(test.url)
      if err != nil {
         t.Fatal(err)
      }
      movie, err := web.Movie()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", movie)
      name, err := text.Name(movie)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
