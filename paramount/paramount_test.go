package paramount

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestItemUs(t *testing.T) {
   var app AppToken
   err := app.ComCbsApp()
   if err != nil {
      t.Fatal(err)
   }
   items, err := app.Items(tests["us"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", items)
}

func TestItemFr(t *testing.T) {
   var app AppToken
   err := app.ComCbsCa()
   if err != nil {
      t.Fatal(err)
   }
   items, err := app.Items(tests["fr"].content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", items)
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
      var pssh widevine.Pssh
      pssh.ContentId = []byte(test.content_id)
      pssh.KeyId, err = hex.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Marshal())
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
