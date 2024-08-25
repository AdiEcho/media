package criterion

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)

func TestVideo(t *testing.T) {
   var (
      token AuthToken
      err error
   )
   token.Raw, err = os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err = token.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   item, err := token.Video(test.slug)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
   name, err := text.Name(item)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", name)
}

var test = struct{
   key_id string
   slug string
   url string
}{
   key_id: "e4576465a745213f336c1ef1bf5d513e",
   slug: "my-dinner-with-andre",
   url: "criterionchannel.com/videos/my-dinner-with-andre",
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
   var token AuthToken
   token.Raw, err = os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err = token.Unmarshal(); err != nil {
      t.Fatal(err)
   }
   item, err := token.Video(test.slug)
   if err != nil {
      t.Fatal(err)
   }
   files, err := token.Files(item)
   if err != nil {
      t.Fatal(err)
   }
   file, ok := files.Dash()
   if !ok {
      t.Fatal("VideoFiles.Dash")
   }
   key, err := module.Key(file, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
