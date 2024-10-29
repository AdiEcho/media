package hulu

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "path"
   "strings"
   "testing"
   "time"
)

func TestAuthenticate(t *testing.T) {
   username, password, ok := strings.Cut(os.Getenv("hulu"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   var auth Authenticate
   err := auth.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.txt", auth.Raw, os.ModePerm)
}

var tests = []struct{
   content string
   key_id string
   url string
}{
   {
      content: "episode",
      key_id: "21b82dc2ebb24d5aa9f8631f04726650",
      url: "hulu.com/watch/023c49bf-6a99-4c67-851c-4c9e7609cc1d",
   },
   {
      content: "film",
      url: "hulu.com/watch/f70dfd4d-dbfb-46b8-abb3-136c841bba11",
   },
}

func TestDetails(t *testing.T) {
   var (
      auth Authenticate
      err error
   )
   auth.Raw, err = os.ReadFile("authenticate.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err = auth.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      base := path.Base(test.url)
      link, err := auth.DeepLink(&EntityId{base})
      if err != nil {
         t.Fatal(err)
      }
      details, err := auth.Details(link)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", details)
      name, err := text.Name(details)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
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
   for _, test := range tests {
      var pssh widevine.Pssh
      pssh.KeyId, err = hex.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      var auth Authenticate
      auth.Raw, err = os.ReadFile("authenticate.txt")
      if err != nil {
         t.Fatal(err)
      }
      if err = auth.Unmarshal(); err != nil {
         t.Fatal(err)
      }
      base := path.Base(test.url)
      link, err := auth.DeepLink(&EntityId{base})
      if err != nil {
         t.Fatal(err)
      }
      play, err := auth.Playlist(link)
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(play, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}
