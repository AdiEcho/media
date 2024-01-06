package nbc

import (
   "154.pages.dev/stream"
   "fmt"
   "testing"
   "time"
)

func Test_Metadata(t *testing.T) {
   for _, mpx_guid := range mpx_guids {
      meta, err := New_Metadata(mpx_guid)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(stream.Name(meta))
      time.Sleep(time.Second)
   }
}

