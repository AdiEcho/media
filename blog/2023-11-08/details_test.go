package hulu

import (
   "encoding/json"
   "os"
   "testing"
)

const eab = "EAB::023c49bf-6a99-4c67-851c-4c9e7609cc1d::196861183::262714326"

func Test_Details(t *testing.T) {
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
   res, err := auth.details(eab)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   os.Stdout.ReadFrom(res.Body)
}
