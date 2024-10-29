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
      t.Fatal("OperationArticle.Film")
   }
   var user OperationUser
   user.Raw, err = os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err = user.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   play, err := user.Play(asset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(play.Dash())
}
