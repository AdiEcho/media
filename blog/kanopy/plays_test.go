package kanopy

import (
   "fmt"
   "reflect"
   "testing"
)

var size_tests = []any{
   video_manifest{},
   video_plays{},
   web_token{},
}

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
