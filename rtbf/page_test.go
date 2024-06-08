package rtbf

import (
   "154.pages.dev/text"
   "fmt"
   "testing"
   "time"
)

func TestPage(t *testing.T) {
   for _, medium := range media {
      page, err := new_page(medium.path)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", page)
      name, err := text.Name(page)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
