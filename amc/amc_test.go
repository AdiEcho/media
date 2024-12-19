package amc

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/base64"
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
      fmt.Printf("%q\n", text.Name(video))
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
func TestLicense(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range key_tests {
      var pssh widevine.Pssh
      pssh.KeyId, err = base64.StdEncoding.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Module
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      data, err := os.ReadFile("amc.txt")
      if err != nil {
         t.Fatal(err)
      }
      var auth Authorization
      err = auth.Unmarshal(data)
      if err != nil {
         t.Fatal(err)
      }
      var web Address
      err = web.Set(test.url)
      if err != nil {
         t.Fatal(err)
      }
      play, err := auth.Playback(web.Nid)
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(play, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}

var key_tests = []struct{
   key_id string
   url string
}{
   {
      key_id: "+7nUc5piRu2GY3lAiA4MvQ==",
      url: "amcplus.com/movies/nocebo--1061554",
   },
   {
      key_id: "vHkdO0RPSsqD3iPzeupPeA==",
      url: "amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152",
   },
}

var path_tests = []string{
   "/movies/nocebo--1061554",
   "amcplus.com/movies/nocebo--1061554",
   "https://www.amcplus.com/movies/nocebo--1061554",
   "www.amcplus.com/movies/nocebo--1061554",
}

func TestPath(t *testing.T) {
   for _, test := range path_tests {
      var web Address
      err := web.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(web)
   }
}
