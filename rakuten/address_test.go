package rakuten

import (
   "41.neocities.org/text"
   "fmt"
   "testing"
   "time"
)

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
      name := text.Name(movie)
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
