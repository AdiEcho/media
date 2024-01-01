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
var test_deep = &Deep_Link{
   "EAB::023c49bf-6a99-4c67-851c-4c9e7609cc1d::196861183::262714326",
}

func user_info() (map[string]string, error) {
   s, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   b, err := os.ReadFile(s + "/hulu/password.json")
   if err != nil {
      return nil, err
   }
   var m map[string]string
   json.Unmarshal(b, &m)
   return m, nil
}
