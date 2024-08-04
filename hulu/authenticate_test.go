package hulu

import (
   "154.pages.dev/text"
   "fmt"
   "os"
   "path"
   "testing"
   "time"
)

func TestDetails(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   raw, err := os.ReadFile(home + "/hulu.txt")
   if err != nil {
      t.Fatal(err)
   }
   var auth Authenticate
   err = auth.Unmarshal(raw)
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      base := path.Base(test.url)
      link, err := auth.DeepLink(EntityId{base})
      if err != nil {
         t.Fatal(err)
      }
      details, err := auth.Details(link)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", details)
      name, err := text.Name(details)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}

var tests = []struct{
   content string
   key_id string
   url string
}{
   {
      content: "episode",
      key_id: "21b82dc2ebb24d5aa9f8631f04726650",
      url: "hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d",
   },
   {
      content: "film",
      url: "hulu.com/watch/f70dfd4d-dbfb-46b8-abb3-136c841bba11",
   },
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
   os.WriteFile("authenticate.txt", auth.Marshal(), 0666)
}
