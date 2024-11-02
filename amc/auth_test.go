package amc

import (
   "41.neocities.org/text"
   "fmt"
   "os"
   "strings"
   "testing"
   "time"
)

func TestContent(t *testing.T) {
   data, err := os.ReadFile("amc.txt")
   if err != nil {
      t.Fatal(err)
   }
   var auth Authorization
   err = auth.Unmarshal(data)
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

func TestRefresh(t *testing.T) {
   data, err := os.ReadFile("amc.txt")
   if err != nil {
      t.Fatal(err)
   }
   var auth Authorization
   err = auth.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   data, err = auth.Refresh()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("amc.txt", data, os.ModePerm)
}

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
   data, err := auth.Login(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("amc.txt", data, os.ModePerm)
}
