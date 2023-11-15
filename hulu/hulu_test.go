package hulu

import (
   "154.pages.dev/http"
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
var watch = Deep_Link{
   "EAB::023c49bf-6a99-4c67-851c-4c9e7609cc1d::196861183::262714326",
}

func user_info() (map[string]string, error) {
   s, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   b, err := os.ReadFile(s + "/hulu.json")
   if err != nil {
      return nil, err
   }
   var m map[string]string
   json.Unmarshal(b, &m)
   return m, nil
}

func Test_Details(t *testing.T) {
   m, err := user_info()
   if err != nil {
      t.Fatal(err)
   }
   http.No_Location()
   http.Verbose()
   auth, err := Living_Room(m["username"], m["password"])
   if err != nil {
      t.Fatal(err)
   }
   detail, err := auth.Details(watch)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", detail)
}

// hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d
const watch_ID = "023c49bf-6a99-4c67-851c-4c9e7609cc1d"

func Test_Deep_Link(t *testing.T) {
   m, err := user_info()
   if err != nil {
      t.Fatal(err)
   }
   http.No_Location()
   http.Verbose()
   auth, err := Living_Room(m["username"], m["password"])
   if err != nil {
      t.Fatal(err)
   }
   link, err := auth.Deep_Link(watch_ID)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", link)
}
