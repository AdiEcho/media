package kanopy

import (
   "fmt"
   "os"
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

func TestPlays(t *testing.T) {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var web web_token
   err = web.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   play, err := web.plays(test.video_id)
   if err != nil {
      t.Fatal(err)
   }
   manifest, ok := play.dash()
   fmt.Printf("%+v %v\n", manifest, ok)
}
