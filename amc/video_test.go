package amc

import (
   "154.pages.dev/http"
   "154.pages.dev/stream"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Login(t *testing.T) {
   auth, err := Unauth()
   if err != nil {
      t.Fatal(err)
   }
   home, err := func() (string, error) {
      s, err := os.UserHomeDir()
      if err != nil {
         return "", err
      }
      return s + "/amc/", nil
   }()
   if err != nil {
      t.Fatal(err)
   }
   u, err := http.User(home + "user.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Login(u.Username, u.Password); err != nil {
      t.Fatal(err)
   }
   text, err := auth.Marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile(home + "auth.json", text, 0666)
}

func Test_Content(t *testing.T) {
   var auth Auth_ID
   {
      s, err := os.UserHomeDir()
      if err != nil {
         t.Fatal(err)
      }
      b, err := os.ReadFile(s + "/amc/auth.json")
      if err != nil {
         t.Fatal(err)
      }
      auth.Unmarshal(b)
   }
   http.No_Location()
   http.Verbose()
   for _, test := range tests {
      con, err := auth.Content(test.address)
      if err != nil {
         t.Fatal(err)
      }
      vid, err := con.Video()
      if err != nil {
         t.Fatal(err)
      }
      name, err := stream.Name(vid)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
}

func Test_Refresh(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var auth Auth_ID
   {
      b, err := os.ReadFile(home + "/amc/auth.json")
      if err != nil {
         t.Fatal(err)
      }
      auth.Unmarshal(b)
   }
   if err := auth.Refresh(); err != nil {
      t.Fatal(err)
   }
   {
      b, err := auth.Marshal()
      if err != nil {
         t.Fatal(err)
      }
      os.WriteFile(home + "/amc/auth.json", b, 0666)
   }
}
