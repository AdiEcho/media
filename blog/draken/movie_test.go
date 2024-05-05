package draken

import (
   "fmt"
   "testing"
   "time"
)

var custom_ids = []string{
   // drakenfilm.se/film/michael-clayton
   "michael-clayton",
   // drakenfilm.se/film/the-card-counter
   "the-card-counter",
}

func TestMovie(t *testing.T) {
   for _, id := range custom_ids {
      movie, err := new_movie(id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", movie)
      time.Sleep(time.Second)
   }
}
