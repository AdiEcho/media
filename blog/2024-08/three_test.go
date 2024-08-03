package http

import "testing"

func BenchmarkThree(b *testing.B) {
   var three response_three
   err := three.New()
   if err != nil {
      b.Fatal(err)
   }
   for range b.N {
      text, err := three.marshal()
      if err != nil {
         b.Fatal(err)
      }
      err = three.unmarshal(text)
      if err != nil {
         b.Fatal(err)
      }
   }
}
