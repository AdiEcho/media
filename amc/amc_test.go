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
      url: "amcplus.com/movies/nocebo--1061554",
   },
   "show": {
      url: "amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152",
      key_id: "bc791d3b444f4aca83de23f37aea4f78",
   },
}

func TestLogin(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   raw, err := Unauth()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := raw.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   raw, err = auth.Login(u["username"], u["password"])
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile(home + "/amc/auth.json", raw, 0666)
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
   test := tests[0]
   key_id, err := hex.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.NewModule(private_key, client_id, key_id, nil)
   if err != nil {
      t.Fatal(err)
   }
   raw, err := os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth, err := RawAuth.Unmarshal(raw)
   if err != nil {
      t.Fatal(err)
   }
   play, err := auth.Playback(test.u)
   if err != nil {
      t.Fatal(err)
   }
   key, err := mod.Key(play)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestRefresh(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   raw, err := os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth, err := RawAuth.Unmarshal(raw)
   if err != nil {
      t.Fatal(err)
   }
   raw, err = auth.Refresh()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile(home + "/amc/auth.json", raw, 0666)
}

var path_tests = []string{
   "http://amcplus.com/movies/nocebo--1061554",
   "amcplus.com/movies/nocebo--1061554",
}

func TestPath(t *testing.T) {
   for _, test := range path_tests {
      var u URL
      err := u.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(u)
   }
}
