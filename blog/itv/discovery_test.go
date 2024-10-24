package itv

import (
   "fmt"
   "testing"
   "time"
)

func TestDiscovery(t *testing.T) {
   for _, test := range tests {
      var title discovery_title
      err := title.New(test.legacy_id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", title)
      time.Sleep(time.Second)
   }
}

var tests = []struct{
   legacy_id string
   url string
}{
   {
      legacy_id: "10/3463/0001",
      url: "itv.com/watch/pulp-fiction/10a3463",
   },
   {
      legacy_id: "10/3915/0002",
      url: "itv.com/watch/community/10a3915/10a3915a0002",
   },
}
