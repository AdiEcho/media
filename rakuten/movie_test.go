package rakuten

import (
   "154.pages.dev/text"
   "fmt"
   "testing"
   "time"
)

var tests = map[string]movie_test{
   "fr": {
      url:        "rakuten.tv/fr/movies/jerry-maguire",
      content_id: "cf34a03bbb4aa899df42c3cf7a6cb925-mc-0-137-0-0",
      key_id:     "00000000000000000000000000000000",
   },
   "se": {
      url:        "rakuten.tv/se/movies/i-heart-huckabees",
      content_id: "9a534a1f12d68e1a2359f38710fddb65-mc-0-147-0-0",
      key_id:     "00000000000000000000000000000000",
   },
}

type movie_test struct {
   content_id string
   key_id     string
   url        string
}

func TestMovie(t *testing.T) {
   for _, test := range tests {
      var web WebAddress
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
