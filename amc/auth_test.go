package amc

import (
   "154.pages.dev/log"
   "154.pages.dev/stream"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

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
   os.WriteFile(home + "/amc/auth.json", auth.Raw, 0666)
}

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
   log.Set_Handler(log.Handler{})
   log.Set_Transport(0)
   var auth Auth_ID
   auth.Raw, err = os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
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

func Test_Content(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var auth Auth_ID
   auth.Raw, err = os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   for _, test := range tests {
      con, err := auth.Content(test.u)
      if err != nil {
         t.Fatal(err)
      }
      vid, err := con.Video()
      if err != nil {
         t.Fatal(err)
      }
      name, err := stream.Format_Film(vid)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
   }
}

func Test_Refresh(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var auth Auth_ID
   auth.Raw, err = os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   if err := auth.Refresh(); err != nil {
      t.Fatal(err)
   }
   os.WriteFile(home + "/amc/auth.json", auth.Raw, 0666)
}

