package kanopy

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "reflect"
   "testing"
)

var test = struct{
   key_id string
   url string
   video_id int
}{
   key_id: "DUCS1DH4TB6Po1oEkG9xUA==",
   url: "kanopy.com/en/product/13808102",
   video_id: 13808102,
}

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
   web_token{},
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
   pssh.KeyId, err = base64.StdEncoding.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      t.Fatal(err)
   }
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var web web_token
   err = web.unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(&web, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
