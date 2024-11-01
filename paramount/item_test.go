package paramount

import (
   "41.neocities.org/text"
   "fmt"
   "reflect"
   "testing"
   "time"
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
   AppToken{},
   Number(0),
   SessionToken{},
   VideoItem{},
}

func TestItemUsa(t *testing.T) {
   var token AppToken
   err := token.ComCbsApp()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      if test.location == "" {
         var item VideoItem
         data, err := item.Marshal(token, test.content_id)
         if err != nil {
            t.Fatal(err)
         }
         err = item.Unmarshal(data)
         if err != nil {
            t.Fatal(err)
         }
         name, err := text.Name(&item)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%q\n", name)
         time.Sleep(time.Second)
      }
   }
}
