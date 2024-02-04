package hulu

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

func Test_Details(t *testing.T) {
   m, err := user_info()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Living_Room(m["username"], m["password"])
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   detail, err := auth.Details(test_deep)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", detail)
}

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
var test_id = ID{"023c49bf-6a99-4c67-851c-4c9e7609cc1d"}

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
var test_deep = &Deep_Link{
   "EAB::023c49bf-6a99-4c67-851c-4c9e7609cc1d::196861183::262714326",
}

func Test_Deep_Link(t *testing.T) {
   m, err := user_info()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Living_Room(m["username"], m["password"])
   if err != nil {
      t.Fatal(err)
   }
   link, err := auth.Deep_Link(test_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", link)
}
