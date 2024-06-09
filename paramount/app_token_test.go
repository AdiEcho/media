package paramount

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "path"
   "testing"
   "time"
)

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
   var token AppToken
   err = token.New()
   if err != nil {
      t.Fatal(err)
   }
   session, err := token.Session(path.Base(test.url))
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(session, pssh.KeyId)
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
      name, err := text.Name(item)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
