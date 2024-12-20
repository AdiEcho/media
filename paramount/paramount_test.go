package paramount

import (
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestWidevine(t *testing.T) {
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
   var app AppToken
   err = app.ComCbsApp()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var pssh widevine.Pssh
      pssh.ContentId = []byte(test.content_id)
      pssh.KeyId, err = base64.StdEncoding.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Module
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      session, err := app.Session(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(session, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}
// need all of these for `assetTypes` test
var tests = []struct{
   content_id string
   location string
   url string
   key_id string
}{
   {
      content_id: "Oo75PgAbcmt9xqqn1AMoBAfo190Cfhqi",
      key_id: "3RyyVzthSSOklAXiQ2vyRw==",
      url: "paramountplus.com/movies/video/Oo75PgAbcmt9xqqn1AMoBAfo190Cfhqi",
   },
   {
      content_id: "esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
      key_id: "H94BVNcqT0WRKzTwzgd36w==",
      url: "paramountplus.com/shows/video/esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
   },
   {
      content_id: "rZ59lcp4i2fU4dAaZJ_iEgKqVg_ogrIf",
      key_id: "Sryog4HeT2CLHx38NftIMA==",
      url: "cbs.com/shows/video/rZ59lcp4i2fU4dAaZJ_iEgKqVg_ogrIf",
   },
   {
      content_id: "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
      key_id: "BsO37qHORXefruKryNAaVQ==",
      location: "France",
      url: "paramountplus.com/movies/video/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
   },
   {
      content_id: "WNujiS5PHkY5wN9doNY6MSo_7G8uBUcX",
      key_id: "bsT01+Q1Ta+39TayayKhBg==",
      location: "Australia",
      url: "paramountplus.com/shows/video/WNujiS5PHkY5wN9doNY6MSo_7G8uBUcX",
   },
}
