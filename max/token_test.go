package max

import (
   "fmt"
   "reflect"
   "testing"
)

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Address{},
   BoltToken{},
   DefaultRoutes{},
   LinkInitiate{},
   LinkLogin{},
   Manifest{},
   Playback{},
   RouteInclude{},
   playback_request{},
}
