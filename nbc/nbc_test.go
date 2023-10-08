package nbc

import (
   "fmt"
   "testing"
   "time"
)

func Test_Video(t *testing.T) {
   for _, guid := range guids {
      meta, err := New_Metadata(guid)
      if err != nil {
         t.Fatal(err)
      }
      on, err := meta.On_Demand()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", on)
      time.Sleep(time.Second)
   }
}
