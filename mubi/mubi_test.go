package mubi

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

// mubi.com/films/325455/player
// mubi.com/films/passages-2022
const passages_2022 = 325455

func TestSecure(t *testing.T) {
   var (
      auth Authenticate
      err error
   )
   auth.Data, err = os.ReadFile("authenticate.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   secure, err := auth.URL(&FilmResponse{ID: passages_2022})
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", secure)
}

// mubi.com/films/325455/player
// mubi.com/films/passages-2022
const default_kid = "CA215A25BB2D43F0BD095FC671C984EE"

func TestLicense(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var auth Authenticate
   auth.Data, err = os.ReadFile(home + "/hulu.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   key_id, err := hex.DecodeString(default_kid)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, widevine.PSSH(key_id))
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(auth, key_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestAuthenticate(t *testing.T) {
   var (
      code LinkCode
      err error
   )
   code.Data, err = os.ReadFile("code.json")
   if err != nil {
      t.Fatal(err)
   }
   code.Unmarshal()
   auth, err := code.Authenticate()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.json", auth.Data, 0666)
}

func TestCode(t *testing.T) {
   var code LinkCode
   err := code.New()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.json", code.Data, 0666)
   code.Unmarshal()
   fmt.Println(code)
}
