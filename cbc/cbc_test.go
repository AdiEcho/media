package gem

import (
   "154.pages.dev/http"
   "fmt"
   "os"
   "testing"
)

func Test_New_Profile(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := http.User(home + "/cbc/user.json")
   if err != nil {
      t.Fatal(err)
   }
   tok, err := New_Token(u.Username, u.Password)
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
   home, err := func() (string, error) {
      s, err := os.UserHomeDir()
      if err != nil {
         return "", err
      }
      return s + "/cbc/", nil
   }()
   if err != nil {
      t.Fatal(err)
   }
   u, err := http.User(home + "user.json")
   if err != nil {
      t.Fatal(err)
   }
   tok, err := New_Token(u.Username, u.Password)
   if err != nil {
      t.Fatal(err)
   }
   pro, err := tok.Profile()
   if err != nil {
      t.Fatal(err)
   }
   if err := pro.Write_File(home + "profile.json"); err != nil {
      t.Fatal(err)
   }
}
