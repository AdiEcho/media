package amc

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "reflect"
   "testing"
   "time"
)

var size_tests = []any{
   &struct{}{},
   Address{},
   Authorization{},
   ContentCompiler{},
   CurrentVideo{},
   Playback{},
}

func TestSize(t *testing.T) {
   for _, test := range size_tests {
      fmt.Println(reflect.TypeOf(test).Size())
   }
}

var path_tests = []string{
   "http://amcplus.com/movies/nocebo--1061554",
   "amcplus.com/movies/nocebo--1061554",
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

var tests = []struct{
   key_id string
   url string
}{
   {
      key_id: "Xn02m57KRCakPhWnbwndfg==",
      url: "amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152",
   },
   {
      url: "amcplus.com/movies/nocebo--1061554",
   },
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
   for _, test := range tests {
      var pssh widevine.Pssh
      pssh.KeyId, err = base64.StdEncoding.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      module.New(private_key, client_id, pssh.Marshal())
      var auth Authorization
      auth.Raw, err = os.ReadFile("/authorization.txt")
      if err != nil {
         t.Fatal(err)
      }
      auth.Unmarshal()
      var web Address
      web.Set(test.url)
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
