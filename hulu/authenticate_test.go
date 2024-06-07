package hulu

import (
   "154.pages.dev/text"
   "fmt"
   "os"
   "testing"
)

var tests = map[string]struct{
   id string
   key_id string
   url string
}{
   "episode": {
      id: "023c49bf-6a99-4c67-851c-4c9e7609cc1d",
      key_id: "21b82dc2ebb24d5aa9f8631f04726650",
      url: "hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d",
   },
}

func TestDetails(t *testing.T) {
   var (
      auth Authenticate
      err error
   )
   auth.Data, err = os.ReadFile("authenticate.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   id := EntityId{tests["episode"].id}
   link, err := auth.DeepLink(id)
   if err != nil {
      t.Fatal(err)
   }
   details, err := auth.Details(link)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", details)
   name, err := text.Name(details[0])
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", name)
}

func TestAuthenticate(t *testing.T) {
   username := os.Getenv("hulu_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("hulu_password")
   var auth Authenticate
   err := auth.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.json", auth.Data, 0666)
}
