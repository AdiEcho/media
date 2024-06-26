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
   var pssh widevine.PSSH
   pssh.KeyId, err = hex.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, pssh.Encode())
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
   key, err := module.Key(play, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestActivationCode(t *testing.T) {
   var token AccountToken
   err := token.New(nil)
   if err != nil {
      t.Fatal(err)
   }
   code, err := token.Code()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(code)
   text, err := code.Marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("code.json", text, 0666)
}

func TestAccountToken(t *testing.T) {
   var (
      activate ActivationToken
      err      error
   )
   activate.Data, err = os.ReadFile("token.json")
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

func TestActivationToken(t *testing.T) {
   text, err := os.ReadFile("code.json")
   if err != nil {
      t.Fatal(err)
   }
   var code ActivationCode
   err = code.Unmarshal(text)
   if err != nil {
      t.Fatal(err)
   }
   token, err := code.Token()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.json", token.Data, 0666)
}

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
