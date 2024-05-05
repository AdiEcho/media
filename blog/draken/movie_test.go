package draken

import (
   "fmt"
   "testing"
   "time"
)

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
