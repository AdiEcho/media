package rakuten

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
)

type movie_test struct {
   content_id string
   key_id     string
   url        string
}

var tests = map[string]movie_test{
   "fr": {
      content_id: "MGU1MTgwMDA2Y2Q1MDhlZWMwMGQ1MzVmZWM2YzQyMGQtbWMtMC0xNDEtMC0w",
      key_id:     "DlGAAGzVCO7ADVNf7GxCDQ==",
      url:        "rakuten.tv/fr/movies/infidele",
   },
   "se": {
      content_id: "OWE1MzRhMWYxMmQ2OGUxYTIzNTlmMzg3MTBmZGRiNjUtbWMtMC0xNDctMC0w",
      key_id:     "mlNKHxLWjhojWfOHEP3bZQ==",
      url:        "rakuten.tv/se/movies/i-heart-huckabees",
   },
}

func (m movie_test) license() ([]byte, error) {
   home, err := os.UserHomeDir()
   if err != nil {
      return nil, err
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      return nil, err
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      return nil, err
   }
   var pssh widevine.Pssh
   pssh.ContentId, err = base64.StdEncoding.DecodeString(m.content_id)
   if err != nil {
      return nil, err
   }
   pssh.KeyId, err = base64.StdEncoding.DecodeString(m.key_id)
   if err != nil {
      return nil, err
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      return nil, err
   }
   var web Address
   web.Set(m.url)
   info, err := web.Hd().Info()
   if err != nil {
      return nil, err
   }
   return module.Key(info, pssh.KeyId)
}

func TestFr(t *testing.T) {
   var web Address
   web.Set(tests["fr"].url)
   stream, err := web.Fhd().Info()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", stream)
}

func TestLicenseFr(t *testing.T) {
   key, err := tests["fr"].license()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestSe(t *testing.T) {
   var web Address
   web.Set(tests["se"].url)
   stream, err := web.Fhd().Info()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", stream)
}

func TestLicenseSe(t *testing.T) {
   key, err := tests["se"].license()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
