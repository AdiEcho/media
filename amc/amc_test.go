package amc

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
)

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
   test := tests["show"]
   var pssh widevine.Pssh
   pssh.KeyId, err = base64.StdEncoding.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Encode())
   if err != nil {
      t.Fatal(err)
   }
   text, err := os.ReadFile(home + "/amc.json")
   if err != nil {
      t.Fatal(err)
   }
   auth, err := RawAuthorization.Authorization(text)
   if err != nil {
      t.Fatal(err)
   }
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
}

var tests = map[string]struct{
   content_id string
   key_id string
   url string
}{
   "movie": {
      url: "amcplus.com/movies/nocebo--1061554",
   },
   "show": {
      key_id: "Xn02m57KRCakPhWnbwndfg==",
      url: "amcplus.com/shows/orphan-black/episodes/season-1-instinct--1011152",
   },
}
