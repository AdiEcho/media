package hulu

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

func Test_Authenticate(t *testing.T) {
   var u struct { Username, Password string }
   {
      s, err := os.UserHomeDir()
      if err != nil {
         t.Fatal(err)
      }
      b, err := os.ReadFile(s + "/hulu.json")
      if err != nil {
         t.Fatal(err)
      }
      json.Unmarshal(b, &u)
   }
   auth, err := living_room(u.Username, u.Password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", auth)
}
