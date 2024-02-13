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
   token, err := NewAppToken()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      item, err := token.Item(path.Base(test.url))
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(rosso.Name(item))
      time.Sleep(time.Second)
   }
}

func TestPost(t *testing.T) {
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
   pssh, err := base64.StdEncoding.DecodeString(test.pssh)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.NewModule(private_key, client_id, nil, pssh)
   if err != nil {
      t.Fatal(err)
   }
   token, err := NewAppToken()
   if err != nil {
      t.Fatal(err)
   }
   sess, err := token.Session(path.Base(test.url))
   if err != nil {
      t.Fatal(err)
   }
   key, err := mod.Key(sess)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
