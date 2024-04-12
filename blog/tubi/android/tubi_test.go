package tubi

import (
   "fmt"
   "testing"
)

// tubitv.com/movies/589292
const content_id = 589292

func TestContent(t *testing.T) {
   var content content_management
   err := content.New(content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", content)
}
