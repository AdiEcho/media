package amc

import (
   "154.pages.dev/text"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestContent(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var auth Authorization
   auth.Data, err = os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   for _, test := range tests {
      var web WebAddress
      web.Set(test.url)
      content, err := auth.Content(web.Path)
      if err != nil {
         t.Fatal(err)
      }
      video, err := content.Video()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(text.Name(video))
      time.Sleep(time.Second)
   }
}

var path_tests = []string{
   "http://amcplus.com/movies/nocebo--1061554",
   "amcplus.com/movies/nocebo--1061554",
}

func TestLogin(t *testing.T) {
   var auth Authorization
   err := auth.Unauth()
   if err != nil {
      t.Fatal(err)
   }
   err = auth.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   username, password := os.Getenv("amc_username"), os.Getenv("amc_password")
   err = auth.Login(username, password)
   if err != nil {
      t.Fatal(err)
   }
}

func TestRefresh(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var auth Authorization
   auth.Data, err = os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   auth.Refresh()
   os.WriteFile(home + "/amc/auth.json", auth.Data, 0666)
}

func TestPath(t *testing.T) {
   for _, test := range path_tests {
      var web WebAddress
      err := web.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(web)
   }
}
