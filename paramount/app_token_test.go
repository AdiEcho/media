package paramount

import (
   "154.pages.dev/media/internal"
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
   if err := token.New(); err != nil {
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
   if err := module.New(private_key, client_id, key_id); err != nil {
      t.Fatal(err)
   }
   license, err := module.License(session)
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Printf("%x %v\n", key, ok)
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
      fmt.Println(internal.Name(<-item))
      time.Sleep(time.Second)
   }
}
