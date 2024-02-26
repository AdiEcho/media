package peacock

import (
   "154.pages.dev/encoding"
   "fmt"
   "testing"
)

func TestQuery(t *testing.T) {
   var node query_node
   err := node.New(content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(encoding.Name(node))
}
