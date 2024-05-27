package member

import (
   "154.pages.dev/text"
   "fmt"
   "testing"
)

// cinemember.nl/films/american-hustle
const american_hustle ArticleSlug = "films/american-hustle"

func TestArticle(t *testing.T) {
   article, err := american_hustle.Article()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", article)
   name, err := text.Name(Encoding{Article: article})
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", name)
}
