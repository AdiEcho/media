package plex

import (
   "fmt"
   "testing"
)

func TestAnonymous(t *testing.T) {
   var anon anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", anon)
}
