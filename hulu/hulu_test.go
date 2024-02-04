package hulu

import (
   "fmt"
   "os"
   "testing"
)

func new_auth() (*Authenticate, error) {
   auth, err := Living_Room(
      os.Getenv("hulu_username"), os.Getenv("hulu_password"),
   )
   if err != nil {
      return nil, err
   }
   if err := auth.Unmarshal(); err != nil {
      return nil, err
   }
   return auth, nil
}

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
var test_id = ID{"023c49bf-6a99-4c67-851c-4c9e7609cc1d"}

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
var test_deep = &Deep_Link{
   "EAB::023c49bf-6a99-4c67-851c-4c9e7609cc1d::196861183::262714326",
}

func Test_Details(t *testing.T) {
   auth, err := new_auth()
   if err != nil {
      t.Fatal(err)
   }
   detail, err := auth.Details(test_deep)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", detail)
}

func Test_Deep_Link(t *testing.T) {
   auth, err := new_auth()
   if err != nil {
      t.Fatal(err)
   }
   link, err := auth.Deep_Link(test_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", link)
}
