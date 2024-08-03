package http

import "testing"

func BenchmarkOne(b *testing.B) {
   var one response_one
   err := one.New()
   if err != nil {
      b.Fatal(err)
   }
   for range b.N {
      err = one.unmarshal()
      if err != nil {
         b.Fatal(err)
      }
   }
}
