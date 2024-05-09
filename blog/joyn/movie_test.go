package joyn

import (
   "fmt"
   "testing"
)

// joyn.de/filme/barry-seal-only-in-america
const barry_seal = "/filme/barry-seal-only-in-america"

func TestMovie(t *testing.T) {
   var movie movie_detail
   err := movie.New(barry_seal)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", movie)
}
