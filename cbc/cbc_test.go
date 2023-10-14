package gem

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

func user(s string) (map[string]string, error) {
   b, err := os.ReadFile(s)
   if err != nil {
      return nil, err
   }
   var m map[string]string
   json.Unmarshal(b, &m)
   return m, nil
}

func Test_New_Profile(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user(home + "/cbc/user.json")
   if err != nil {
      t.Fatal(err)
   }
   tok, err := New_Token(u["username"], u["password"])
   if err != nil {
      t.Fatal(err)
   }
   pro, err := tok.Profile()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", pro)
}

func Test_Profile(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user(home + "/cbc/user.json")
   if err != nil {
      t.Fatal(err)
   }
   tok, err := New_Token(u["username"], u["password"])
   if err != nil {
      t.Fatal(err)
   }
   pro, err := tok.Profile()
   if err != nil {
      t.Fatal(err)
   }
   if err := pro.Write_File(home + "/cbc/profile.json"); err != nil {
      t.Fatal(err)
   }
}
