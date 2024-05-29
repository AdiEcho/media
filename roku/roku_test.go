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

func TestPlayback(t *testing.T) {
   var token AccountToken
   err := token.New(nil)
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      play, err := token.Playback(path.Base(test.url))
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", play)
      time.Sleep(time.Second)
   }
}

func TestActivationToken(t *testing.T) {
   text, err := os.ReadFile("code.json")
   if err != nil {
      t.Fatal(err)
   }
   var code activation_code
   err = code.unmarshal(text)
   if err != nil {
      t.Fatal(err)
   }
   token, err := code.token()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.json", token.data, 0666)
}

func TestAccountToken(t *testing.T) {
   var (
      activate activation_token
      err error
   )
   activate.data, err = os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   err = activate.unmarshal()
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
func TestLicense(t *testing.T) {
   test := tests["episode"]
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
   key_id, err := hex.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, widevine.PSSH(key_id, nil))
   if err != nil {
      t.Fatal(err)
   }
   var token AccountToken
   err = token.New(nil)
   if err != nil {
      t.Fatal(err)
   }
   play, err := token.Playback(path.Base(test.url))
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(play, key_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
