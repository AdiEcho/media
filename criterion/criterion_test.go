package criterion

import (
   "154.pages.dev/text"
   "fmt"
   "os"
   "testing"
)

func TestToken(t *testing.T) {
   username := os.Getenv("criterion_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("criterion_password")
   data, err := NewAuthToken(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.json", data, 0666)
}

func TestVideo(t *testing.T) {
   data, err := os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   var token AuthToken
   err = token.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   item, err := token.Video(my_dinner)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
   name, err := text.Name(item)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", name)
}

// criterionchannel.com/videos/my-dinner-with-andre
const my_dinner = "my-dinner-with-andre"
