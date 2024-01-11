package amc

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

var tests = []struct {
   key_ID string
   u URL
} {
   { // amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152
      key_ID: "bc791d3b444f4aca83de23f37aea4f78",
      u: URL{"/shows/orphan-black/episodes/season-1-instinct--1011152", "1011152"},
   },
   { // amcplus.com/movies/nocebo--1061554
      u: URL{"/movies/nocebo--1061554", "1061554"},
   },
}

func Test_Login(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user(home + "/amc/user.json")
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

func Test_Key(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_ID, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   test := tests[0]
   key_ID, err := hex.DecodeString(test.key_ID)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, key_ID, nil)
   if err != nil {
      t.Fatal(err)
   }
   raw, err := os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Raw_Auth.Unmarshal(raw)
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

func Test_Refresh(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   raw, err := os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Raw_Auth.Unmarshal(raw)
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

func Test_Path(t *testing.T) {
   for _, test := range path_tests {
      var u URL
      err := u.Set(test)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(u)
   }
}

func user(s string) (map[string]string, error) {
   b, err := os.ReadFile(s)
   if err != nil {
      return nil, err
   }
   var m map[string]string
   json.Unmarshal(b, &m)
   return m, nil
}
