package rakuten

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
)

func TestSe(t *testing.T) {
   var video on_demand
   video.fhd(classification["se"], "i-heart-huckabees")
   stream, err := video.stream()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", stream)
}

func TestFr(t *testing.T) {
   var video on_demand
   video.fhd(classification["fr"], "jerry-maguire")
   stream, err := video.stream()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", stream)
}

func TestLicenseFr(t *testing.T) {
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
   test := tests["fr"]
   content_id, err := base64.StdEncoding.DecodeString(test.content_id)
   if err != nil {
      t.Fatal(err)
   }
   key_id, err := base64.StdEncoding.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, widevine.PSSH(key_id, content_id))
   if err != nil {
      t.Fatal(err)
   }
   var video on_demand
   video.hd(classification["fr"], "jerry-maguire")
   stream, err := video.stream()
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(stream, key_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x:%x\n", key_id, key)
}

func TestLicenseSe(t *testing.T) {
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
   test := tests["se"]
   content_id, err := base64.StdEncoding.DecodeString(test.content_id)
   if err != nil {
      t.Fatal(err)
   }
   key_id, err := base64.StdEncoding.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, widevine.PSSH(key_id, content_id))
   if err != nil {
      t.Fatal(err)
   }
   var video on_demand
   video.hd(classification["se"], "i-heart-huckabees")
   stream, err := video.stream()
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(stream, key_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x:%x\n", key_id, key)
}

var tests = map[string]struct{
   content_id string
   key_id string
   url string
}{
   "fr": {
      content_id: "Y2YzNGEwM2JiYjRhYTg5OWRmNDJjM2NmN2E2Y2I5MjUtbWMtMC0xMzctMC0w",
      key_id: "zzSgO7tKqJnfQsPPemy5JQ==",
      url: "rakuten.tv/fr/movies/jerry-maguire",
   },
   "se": {
      content_id: "OWE1MzRhMWYxMmQ2OGUxYTIzNTlmMzg3MTBmZGRiNjUtbWMtMC0xNDctMC0w",
      key_id: "mlNKHxLWjhojWfOHEP3bZQ==",
      url: "rakuten.tv/se/movies/i-heart-huckabees",
   },
}
