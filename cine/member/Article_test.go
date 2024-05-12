package member

import (
   "154.pages.dev/encoding"
   "fmt"
   "testing"
)

const american_hustle = "films/american-hustle"

func TestArticle(t *testing.T) {
   article, err := new_article(american_hustle)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", article)
   name, err := encoding.Name(namer{article})
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", name)
}
