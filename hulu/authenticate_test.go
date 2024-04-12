package hulu

import (
   "fmt"
   "os"
   "testing"
)

func (a *Authenticate) getenv() error {
   err := a.New(os.Getenv("hulu_username"), os.Getenv("hulu_password"))
   if err != nil {
      return err
   }
   return a.Unmarshal()
}

func TestDetails(t *testing.T) {
   var auth Authenticate
   err := auth.getenv()
   if err != nil {
      t.Fatal(err)
   }
   detail, err := auth.Details(test_deep)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", detail)
}

func TestDeepLink(t *testing.T) {
   var auth Authenticate
   err := auth.getenv()
   if err != nil {
      t.Fatal(err)
   }
   link, err := auth.DeepLink(test_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", link)
}

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
var test_id = ID{"023c49bf-6a99-4c67-851c-4c9e7609cc1d"}
