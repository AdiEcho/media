package draken

import (
   "fmt"
   "testing"
)

func TestMovie(t *testing.T) {
   var movie full_movie
   err := movie.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", movie)
}
