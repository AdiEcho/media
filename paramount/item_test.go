package paramount

import (
   "154.pages.dev/http/option"
   "fmt"
   "testing"
   "time"
)

func Test_Item(t *testing.T) {
   option.No_Location()
   option.Verbose()
   token, err := New_App_Token()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      item, err := token.Item(test.content_ID)
      if err != nil {
         t.Fatal(err)
      }
      name, err := media.Name(item)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
}
