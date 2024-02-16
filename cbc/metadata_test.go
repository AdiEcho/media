package cbc

import (
   "154.pages.dev/rosso"
   "fmt"
   "os"
   "testing"
   "time"
)

var links = []string{
   "https://gem.cbc.ca/downton-abbey/s01e05",
   "https://gem.cbc.ca/the-fall/s02e03",
   "https://gem.cbc.ca/the-witch",
}

func TestStream(t *testing.T) {
   for _, link := range links {
      gem, err := NewCatalogGem(link)
      if err != nil {
         t.Fatal(err)
      }
      item := gem.Item()
      fmt.Printf("%+v\n", item)
      fmt.Println(rosso.Name(gem.StructuredMetadata))
      time.Sleep(time.Second)
   }
}

func TestMedia(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   pro, err := ReadProfile(home + "/cbc/profile.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, link := range links {
      gem, err := NewCatalogGem(link)
      if err != nil {
         t.Fatal(err)
      }
      media, err := pro.Media(gem.Item())
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", media)
      time.Sleep(time.Second)
   }
}
