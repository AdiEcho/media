package paramount

import (
   "154.pages.dev/rosso"
   "fmt"
   "testing"
   "time"
)

func Test_Item(t *testing.T) {
   token, err := New_App_Token()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      item, err := token.Item(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(stream.Name(item))
      time.Sleep(time.Second)
   }
}
