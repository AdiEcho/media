package amc

import (
   "154.pages.dev/http"
   "154.pages.dev/widevine"
   "encoding/base64"
   "encoding/json"
   "fmt"
   "os"
   "testing"
)

func Test_Refresh(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var auth Auth_ID
   {
      b, err := os.ReadFile(home + "/amc/auth.json")
      if err != nil {
         t.Fatal(err)
      }
      auth.Unmarshal(b)
   }
   if err := auth.Refresh(); err != nil {
      t.Fatal(err)
   }
   {
      b, err := auth.Marshal()
      if err != nil {
         t.Fatal(err)
      }
      os.WriteFile(home + "/amc/auth.json", b, 0666)
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

func Test_Login(t *testing.T) {
   auth, err := Unauth()
   if err != nil {
      t.Fatal(err)
   }
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   u, err := user(home + "/amc/user.json")
   if err != nil {
      t.Fatal(err)
   }
   if err := auth.Login(u["username"], u["password"]); err != nil {
      t.Fatal(err)
   }
   text, err := auth.Marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile(home + "/amc/auth.json", text, 0666)
}

var tests = []struct {
   address string
   pssh string
} {
   // amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152
   episode: {
      address: "/shows/orphan-black/episodes/season-1-instinct--1011152",
      pssh: "AAAAVnBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADYIARIQuC5UBJ1cQS2w6wxWli1eSxoNd2lkZXZpbmVfdGVzdCIIMTIzNDU2NzgyB2RlZmF1bHQ=",
   },
   // amcplus.com/movies/nocebo--1061554
   movie: {address: "/movies/nocebo--1061554"},
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
   test := tests[episode]
   pssh, err := base64.StdEncoding.DecodeString(test.pssh)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.New_Module(private_key, client_ID, nil, pssh)
   if err != nil {
      t.Fatal(err)
   }
   var auth Auth_ID
   {
      b, err := os.ReadFile(home + "/amc/auth.json")
      if err != nil {
         t.Fatal(err)
      }
      auth.Unmarshal(b)
   }
   http.No_Location()
   http.Verbose()
   play, err := auth.Playback(test.address)
   if err != nil {
      t.Fatal(err)
   }
   key, err := mod.Key(play)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

const (
   episode = iota
   movie
)
