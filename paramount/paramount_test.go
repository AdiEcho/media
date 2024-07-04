package paramount

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestMpdUs(t *testing.T) {
   address, err := MpegDash(tests["us"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(address)
}

func TestMpdFr(t *testing.T) {
   address, err := MpegDash(tests["fr"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(address)
}

func TestWidevine(t *testing.T) {
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
      var pssh widevine.PSSH
      pssh.ContentId = []byte(test.content_id)
      pssh.KeyId, err = hex.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.CDM
      err = module.New(private_key, client_id, pssh.Encode())
      if err != nil {
         t.Fatal(err)
      }
      var app AppToken
      app.ComCbsApp()
      session, err := app.Session(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(session, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}
