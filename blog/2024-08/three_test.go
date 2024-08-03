package http

import (
   "fmt"
   "testing"
)

func TestThree(b *testing.T) {
   var three response_three
   err := three.New()
   if err != nil {
      b.Fatal(err)
   }
   text, err := three.marshal()
   if err != nil {
      b.Fatal(err)
   }
   err = three.unmarshal(text)
   if err != nil {
      b.Fatal(err)
   }
   fmt.Printf("%+v\n", three)
}

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
