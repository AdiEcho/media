package amc

import (
   "154.pages.dev/text"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestContent(t *testing.T) {
   var (
      auth Authorization
      err error
   )
   auth.Raw, err = os.ReadFile("/authorization.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = auth.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var web Address
      web.Set(test.url)
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
   auth.Raw, err = os.ReadFile("/authorization.txt")
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
   os.WriteFile("/authorization.txt", auth.Raw, os.ModePerm)
}
