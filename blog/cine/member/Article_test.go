package member

import (
   "fmt"
   "testing"
)

const american_hustle = "films/american-hustle"

func TestArticle(t *testing.T) {
   var article article_query
   err := article.New(american_hustle)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", article)
}
