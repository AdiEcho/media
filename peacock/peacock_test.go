package peacock

import (
   "154.pages.dev/encoding"
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
)

func TestVideo(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   text, err := os.ReadFile(home + "/peacock.json")
   if err != nil {
      t.Fatal(err)
   }
   var sign SignIn
   sign.Unmarshal(text)
   auth, err := sign.Auth()
   if err != nil {
      t.Fatal(err)
   }
   video, err := auth.Video(content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
}

// peacocktv.com/watch/playback/vod/GMO_00000000224510_02_HDSDR
const (
   content_id = "GMO_00000000224510_02_HDSDR"
   pssh = "AAAAOHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAABgSEAAW4jRz6+d9k9jRpy3GkNdI49yVmwY="
)

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
   data, err := base64.StdEncoding.DecodeString(pssh)
   if err != nil {
      t.Fatal(err)
   }
   var protect widevine.PSSH
   if err := protect.New(data); err != nil {
      t.Fatal(err)
   }
   module, err := protect.CDM(private_key, client_id)
   if err != nil {
      t.Fatal(err)
   }
   data, err = os.ReadFile(home + "/peacock.json")
   if err != nil {
      t.Fatal(err)
   }
   var sign SignIn
   sign.Unmarshal(data)
   auth, err := sign.Auth()
   if err != nil {
      t.Fatal(err)
   }
   video, err := auth.Video(content_id)
   if err != nil {
      t.Fatal(err)
   }
   license, err := module.License(video)
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Println(key, ok)
}

func TestQuery(t *testing.T) {
   var node QueryNode
   err := node.New(content_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(encoding.Name(node))
}
