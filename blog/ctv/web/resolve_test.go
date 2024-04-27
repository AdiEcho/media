package ctv

import (
   "fmt"
   "os"
   "testing"
)

// ctv.ca/movies/the-girl-with-the-dragon-tattoo-2011
const dragon_tattoo = "/movies/the-girl-with-the-dragon-tattoo-2011"

func TestWrite(t *testing.T) {
   play, err := first_playable_content(dragon_tattoo)
   if err != nil {
      t.Fatal(err)
   }
   items, err := play.items()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("contentPackages.json", items.data, 0666)
}

func TestRead(t *testing.T) {
   var (
      items content_packages
      err error
   )
   items.data, err = os.ReadFile("contentPackages.json")
   if err != nil {
      t.Fatal(err)
   }
   item, err := items.item()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
}
