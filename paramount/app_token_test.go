package paramount

import (
   "154.pages.dev/encoding"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "path"
   "testing"
   "time"
)

func TestWidevine(t *testing.T) {
   test := tests["episode"]
   var token AppToken
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   session, err := token.Session(path.Base(test.url))
   if err != nil {
      t.Fatal(err)
   }
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
   err = module.New(private_key, client_id, widevine.PSSH(key_id))
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(session, key_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestItem(t *testing.T) {
   var token AppToken
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      item, err := token.Item(path.Base(test.url))
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(encoding.Name(<-item))
      time.Sleep(time.Second)
   }
}
