package hulu

import (
   "154.pages.dev/text"
   "fmt"
   "os"
   "testing"
   "time"
)

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
   for _, test := range tests {
      link, err := auth.DeepLink(EntityId{test.id})
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
      time.Sleep(time.Second)
   }
}

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
   "film": {
      id: "f70dfd4d-dbfb-46b8-abb3-136c841bba11",
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
   os.WriteFile("authenticate.json", auth.Data, 0666)
}
