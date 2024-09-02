package rakuten

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = map[string]movie_test{
   "se": {
      url:        "rakuten.tv/se/movies/i-heart-huckabees",
      content_id: "9a534a1f12d68e1a2359f38710fddb65-mc-0-147-0-0",
      key_id:     "00000000000000000000000000000000",
   },
}

type movie_test struct {
   content_id string
   key_id     string
   url        string
}

func TestMovie(t *testing.T) {
   for _, test := range tests {
      var web Address
      err := web.Set(test.url)
      if err != nil {
         t.Fatal(err)
      }
      movie, err := web.Movie()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", movie)
      name, err := text.Name(movie)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
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
   pssh.KeyId, err = hex.DecodeString(m.key_id)
   if err != nil {
      return nil, err
   }
   pssh.ContentId = []byte(m.content_id)
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

func TestLicenseSe(t *testing.T) {
   key, err := tests["se"].license()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
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
