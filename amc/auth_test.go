package amc

import (
   "41.neocities.org/text"
   "fmt"
   "os"
   "strings"
   "testing"
   "time"
)

func TestLogin(t *testing.T) {
   username, password, ok := strings.Cut(os.Getenv("amc"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
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
   os.WriteFile("amc.txt", auth.Raw, os.ModePerm)
}

func TestRefresh(t *testing.T) {
   var (
      auth Authorization
      err error
   )
   auth.Raw, err = os.ReadFile("amc.txt")
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
   os.WriteFile("amc.txt", auth.Raw, os.ModePerm)
}

func TestContent(t *testing.T) {
   var (
      auth Authorization
      err error
   )
   auth.Raw, err = os.ReadFile("amc.txt")
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
         t.Fatal("ContentCompiler.Video")
      }
      name, err := text.Name(video)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
