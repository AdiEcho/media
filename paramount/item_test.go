package paramount

import (
   "154.pages.dev/rosso"
   "fmt"
   "testing"
   "time"
)

func TestItem(t *testing.T) {
   token, err := NewAppToken()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      item, err := token.Item(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(rosso.Name(item))
      time.Sleep(time.Second)
   }
}
