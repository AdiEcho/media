package roku

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "path"
   "testing"
   "time"
)

func TestActivationCode(t *testing.T) {
   // AccountToken
   var account AccountToken
   err := account.New(nil)
   if err != nil {
      t.Fatal(err)
   }
   
   
   
   account.Unmarshal()
   // ActivationCode
   code, err := account.Code()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.json", code.Data, 0666)
   code.Unmarshal()
   fmt.Println(code)
}

func TestActivationToken(t *testing.T) {
   var (
      code ActivationCode
      err error
   )
   code.Data, err = os.ReadFile("code.json")
   if err != nil {
      t.Fatal(err)
   }
   err = code.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   
   activation_token, err := code.Token()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("activation_token.json", activation_token.Data, 0666)
}

func TestPlayback(t *testing.T) {
   var account AccountToken
   err := account.New(nil)
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      play, err := account.Playback(path.Base(test.url))
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", play)
      time.Sleep(time.Second)
   }
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
   test := tests["episode"]
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
   var account AccountToken
   err = account.New(nil)
   if err != nil {
      t.Fatal(err)
   }
   play, err := account.Playback(path.Base(test.url))
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(play, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestAccountToken(t *testing.T) {
   var (
      activate ActivationToken
      err      error
   )
   activate.Data, err = os.ReadFile("activate.json")
   if err != nil {
      t.Fatal(err)
   }
   err = activate.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   var account AccountToken
   err = account.New(&activate)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", account)
}
