package kanopy

import (
   "fmt"
   "os"
   "reflect"
   "testing"
)

func TestPlays(t *testing.T) {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token web_token
   err = token.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   plays, err := token.plays(test.video_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", plays)
}

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
