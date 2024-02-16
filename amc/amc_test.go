package amc

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

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

func TestKey(t *testing.T) {
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
   test := tests["show"]
   var protect widevine.PSSH
   protect.Key_ID, err = hex.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   module, err := protect.CDM(private_key, client_id)
   if err != nil {
      t.Fatal(err)
   }
   var auth Authorization
   auth.Raw, err = os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   var web WebAddress
   web.Set(test.url)
   play, err := auth.Playback(web.NID)
   if err != nil {
      t.Fatal(err)
   }
   license, err := module.License(play)
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Printf("%x %v\n", key, ok)
}

var tests = map[string]struct{
   key_id string
   url string
}{
   "movie": {
      url: "http://amcplus.com/movies/nocebo--1061554",
   },
   "show": {
      url: "http://amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152",
      key_id: "bc791d3b444f4aca83de23f37aea4f78",
   },
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
   if err := auth.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   username, password := os.Getenv("amc_username"), os.Getenv("amc_password")
   if err := auth.Login(username, password); err != nil {
      t.Fatal(err)
   }
}

func TestRefresh(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var auth Authorization
   auth.Raw, err = os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   auth.Refresh()
   os.WriteFile(home + "/amc/auth.json", auth.Raw, 0666)
}

