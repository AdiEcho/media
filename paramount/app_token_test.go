package paramount

import (
   "154.pages.dev/rosso"
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "path"
   "testing"
   "time"
)

func TestItem(t *testing.T) {
   var at AppToken
   err := at.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      item, err := at.Item(path.Base(test.url))
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(rosso.Name(item))
      time.Sleep(time.Second)
   }
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
   test := tests["episode cenc"]
   var protect widevine.PSSH
   {
      b, err := base64.StdEncoding.DecodeString(test.pssh)
      if err != nil {
         t.Fatal(err)
      }
      if err := protect.New(b); err != nil {
         t.Fatal(err)
      }
   }
   module, err := protect.CDM(private_key, client_id)
   if err != nil {
      t.Fatal(err)
   }
   var at AppToken
   if err := at.New(); err != nil {
      t.Fatal(err)
   }
   session, err := at.Session(path.Base(test.url))
   if err != nil {
      t.Fatal(err)
   }
   license, err := module.License(session)
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Printf("%x %v\n", key, ok)
}
