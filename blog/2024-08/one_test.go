package http

import (
   "fmt"
   "testing"
)

func TestOne(b *testing.T) {
   var one response_one
   err := one.New()
   if err != nil {
      b.Fatal(err)
   }
   err = one.unmarshal()
   if err != nil {
      b.Fatal(err)
   }
   fmt.Println(one.date.value)
   fmt.Printf("%+v\n", one.body.value)
}

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
