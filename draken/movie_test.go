package draken

import (
   "154.pages.dev/encoding"
   "fmt"
   "path"
   "testing"
   "time"
)

func TestMovie(t *testing.T) {
   for _, film := range films {
      movie, err := NewMovie(path.Base(film.url))
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", movie)
      name, err := encoding.Name(Namer{movie})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(99 * time.Millisecond)
   }
}

var films = []struct {
   content_id string
   key_id     string
   url        string
}{
   {
      content_id: "ODE0OTQ1NWMtY2IzZC00YjE1LTg1YTgtYjk1ZTNkMTU3MGI1",
      key_id:     "e5WypDjIM1+4W74cf6rHIw==",
      url:        "drakenfilm.se/film/michael-clayton",
   },
   {
      content_id: "MTcxMzkzNTctZWQwYi00YTE2LThiZTYtNjllNDE4YzRiYTQw",
      key_id:     "ToV4wH2nlVZE8QYLmLywDg==",
      url:        "drakenfilm.se/film/the-card-counter",
   },
}
