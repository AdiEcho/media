package member

import (
   "fmt"
   "os"
   "testing"
)

func TestAsset(t *testing.T) {
   article, err := american_hustle.Article()
   if err != nil {
      t.Fatal(err)
   }
   asset, ok := article.Film()
   if !ok {
      t.Fatal(ArticleAsset{})
   }
   raw, err := os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   var user OperationUser
   err = user.Unmarshal(raw)
   if err != nil {
      t.Fatal(err)
   }
   play, err := user.Play(asset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(play.Dash())
}
