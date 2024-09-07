package hulu

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "path"
   "reflect"
   "testing"
   "time"
)

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Authenticate{},
   DeepLink{},
   Details{},
   EntityId{},
   Playlist{},
   codec_value{},
   drm_value{},
   playlist_request{},
   segment_value{},
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

func TestAuthenticate(t *testing.T) {
   username := os.Getenv("hulu_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("hulu_password")
   var auth Authenticate
   err := auth.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("authenticate.txt", auth.Raw, os.ModePerm)
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
