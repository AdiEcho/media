package mubi

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

var test = struct{
   id int64
   key_id string
   url []string
}{
   id: 325455,
   key_id: "CA215A25BB2D43F0BD095FC671C984EE",
   url: []string{
      "mubi.com/films/325455/player",
      "mubi.com/films/passages-2022",
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
   var pssh widevine.Pssh
   pssh.KeyId, err = hex.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Encode())
   if err != nil {
      t.Fatal(err)
   }
   var auth Authenticate
   auth.Data, err = os.ReadFile(home + "/authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   key, err := module.Key(auth, pssh.KeyId)
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
   code.Data, err = os.ReadFile("code.txt")
   if err != nil {
      t.Fatal(err)
   }
   code.Unmarshal()
   auth, err := code.Authenticate()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.txt", auth.Data, 0666)
}

func TestCode(t *testing.T) {
   var code LinkCode
   err := code.New()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.txt", code.Data, 0666)
   code.Unmarshal()
   fmt.Println(code)
}

func TestSecure(t *testing.T) {
   var (
      auth Authenticate
      err error
   )
   auth.Data, err = os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   secure, err := auth.Url(&FilmResponse{Id: test.id})
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", secure)
}
