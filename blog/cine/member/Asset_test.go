package member

import (
   "fmt"
   "os"
   "testing"
)

func TestAsset(t *testing.T) {
   article, err := new_article(american_hustle)
   if err != nil {
      t.Fatal(err)
   }
   var auth authenticate
   auth.data, err = os.ReadFile("authenticate.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.unmarshal()
   asset, ok := article.film()
   if !ok {
      t.Fatal("data_article.film")
   }
   play, err := auth.play(asset)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", play)
}
