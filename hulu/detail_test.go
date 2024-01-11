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
