package member

import (
   "154.pages.dev/text"
   "fmt"
   "testing"
)

// cinemember.nl/films/american-hustle
var american_hustle = Address{"films/american-hustle"}

func TestArticle(t *testing.T) {
   article, err := american_hustle.Article()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", article)
   name, err := text.Name(article)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", name)
}
