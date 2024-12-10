package criterion

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "strings"
   "testing"
)

func TestToken(t *testing.T) {
   username, password, ok := strings.Cut(os.Getenv("criterion"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   data, err := (*AuthToken).Marshal(nil, username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", data, os.ModePerm)
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
   pssh.KeyId, err = hex.DecodeString(video_test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Module
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      t.Fatal(err)
   }
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token AuthToken
   err = token.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   item, err := token.Video(video_test.slug)
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

func TestVideo(t *testing.T) {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token AuthToken
   err = token.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   item, err := token.Video(video_test.slug)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", item)
   fmt.Printf("%q\n", text.Name(item))
}

var video_test = struct{
   key_id string
   slug string
   url string
}{
   key_id: "e4576465a745213f336c1ef1bf5d513e",
   slug: "my-dinner-with-andre",
   url: "criterionchannel.com/videos/my-dinner-with-andre",
}
