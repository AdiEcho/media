package joyn

import (
   "fmt"
   "testing"
)

func TestEntitlement(t *testing.T) {
   var anon anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   var movie movie_detail
   err = movie.New(barry_seal)
   if err != nil {
      t.Fatal(err)
   }
   title, err := anon.entitlement(movie)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", title)
}
