package amc

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

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

func TestLicense(t *testing.T) {
   test := tests["show"]
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
   var web WebAddress
   web.Set(test.url)
   play, err := auth.Playback(web.NID)
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
   key_id, err := hex.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   if err := module.New(private_key, client_id, key_id); err != nil {
      t.Fatal(err)
   }
   license, err := module.License(play)
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Printf("%x %v\n", key, ok)
}
