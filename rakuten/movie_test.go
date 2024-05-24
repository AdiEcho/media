package rakuten

import (
   "154.pages.dev/encoding"
   "fmt"
   "testing"
)

func TestMovie(t *testing.T) {
   var movie gizmo_movie
   err := movie.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", movie)
   name, err := encoding.Name(movie)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", name)
}
