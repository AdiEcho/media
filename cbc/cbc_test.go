package cbc

import (
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

func Test_New_Profile(t *testing.T) {
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
