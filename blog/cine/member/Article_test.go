package member

import (
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
}
