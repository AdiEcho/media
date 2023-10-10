package gem

import (
   "154.pages.dev/stream"
   "encoding/json"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Gem(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   pro, err := Read_Profile(home + "/cbc/profile.json")
   if err != nil {
      t.Fatal(err)
   }
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   for _, link := range links {
      gem, err := New_Catalog_Gem(link)
      if err != nil {
         t.Fatal(err)
      }
      item := gem.Item()
      enc.Encode(item)
      name, err := stream.Name(gem.Structured_Metadata)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      media, err := pro.Media(item)
      if err != nil {
         t.Fatal(err)
      }
      enc.Encode(media)
      time.Sleep(time.Second)
   }
}

var links = []string{
   "https://gem.cbc.ca/downton-abbey/s01e05",
   "https://gem.cbc.ca/the-fall/s02e03",
   "https://gem.cbc.ca/the-witch",
}
