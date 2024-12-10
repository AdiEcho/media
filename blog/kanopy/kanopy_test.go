package kanopy

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "strings"
   "testing"
   "time"
)

func TestLogin(t *testing.T) {
   email, password, ok := strings.Cut(os.Getenv("kanopy"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   data, err := web_token{}.marshal(email, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", data, os.ModePerm)
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
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token web_token
   err = token.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   member, err := token.membership()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var pssh widevine.Pssh
      pssh.KeyId, err = base64.StdEncoding.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Module
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      plays, err := token.plays(member, test.video_id)
      if err != nil {
         t.Fatal(err)
      }
      manifest, ok := plays.dash()
      if !ok {
         t.Fatal("video_plays.dash")
      }
      key, err := module.Key(&poster{manifest, &token}, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}
