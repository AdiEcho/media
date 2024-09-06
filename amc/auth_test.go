package amc

import (
   "154.pages.dev/text"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestLogin(t *testing.T) {
   username := os.Getenv("amc_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("amc_password")
   var auth Authorization
   err := auth.Unauth()
   if err != nil {
      t.Fatal(err)
   }
   err = auth.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   err = auth.Login(username, password)
   if err != nil {
      t.Fatal(err)
   }
}

func TestRefresh(t *testing.T) {
   var (
      auth Authorization
      err error
   )
   auth.Raw, err = os.ReadFile("authorization.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = auth.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   err = auth.Refresh()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authorization.txt", auth.Raw, os.ModePerm)
}

func TestContent(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var auth Authorization
   auth.Raw, err = os.ReadFile(home + "/amc.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = auth.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range key_tests {
      var web Address
      err = web.Set(test.url)
      if err != nil {
         t.Fatal(err)
      }
      content, err := auth.Content(web.Path)
      if err != nil {
         t.Fatal(err)
      }
      video, ok := content.Video()
      if !ok {
         t.Fatal(content.VideoError())
      }
      name, err := text.Name(video)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
