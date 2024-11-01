package member

import (
   "fmt"
   "os"
   "testing"
)

func TestAsset(t *testing.T) {
   var article OperationArticle
   data, err := article.Marshal(&american_hustle)
   if err != nil {
      t.Fatal(err)
   }
   err = article.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   asset, ok := article.Film()
   if !ok {
      t.Fatal("OperationArticle.Film")
   }
   data, err = os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   var user OperationUser
   err = user.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   var play OperationPlay
   data, err = play.Marshal(user, asset)
   if err != nil {
      t.Fatal(err)
   }
   err = play.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(play.Dash())
}
