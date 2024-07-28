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
   var user OperationUser
   user.raw, err = os.ReadFile("authenticate.json")
   if err != nil {
      t.Fatal(err)
   }
   err = user.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   play, err := user.Play(asset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(play.Dash())
}
