package rakuten

import (
   "154.pages.dev/text"
   "fmt"
   "testing"
   "time"
)

type movie_test struct {
   content_id string
   key_id     string
   url        string
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
