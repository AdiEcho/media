package paramount

import (
   "41.neocities.org/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

// need all of these for `assetTypes` test
var tests = []struct{
   content_id string
   location string
   url string
   pssh string
}{
   {
      content_id: "Oo75PgAbcmt9xqqn1AMoBAfo190Cfhqi",
      location: "USA",
      pssh: "CAESEN0cslc7YUkjpJQF4kNr8kciIE9vNzVQZ0FiY210OXhxcW4xQU1vQkFmbzE5MENmaHFpOAE=",
      url: "paramountplus.com/movies/video/Oo75PgAbcmt9xqqn1AMoBAfo190Cfhqi",
   },
   {
      content_id: "esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
      location: "USA",
      pssh: "CAESEB/eAVTXKk9FkSs08M4Hd+siIGVzSnZGbHFkcmNTX2tGSG5weFN1WXA0NDlFN3RUZXhEOAE=",
      url: "paramountplus.com/shows/video/esJvFlqdrcS_kFHnpxSuYp449E7tTexD",
   },
   {
      content_id: "rZ59lcp4i2fU4dAaZJ_iEgKqVg_ogrIf",
      location: "USA",
      pssh: "CAESEEq8qIOB3k9gix8d/DX7SDAiIHJaNTlsY3A0aTJmVTRkQWFaSl9pRWdLcVZnX29ncklmOAE=",
      url: "cbs.com/shows/video/rZ59lcp4i2fU4dAaZJ_iEgKqVg_ogrIf",
   },
   {
      content_id: "Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
      location: "France",
      pssh: "CAESEAbDt+6hzkV3n67iq8jQGlUiIFk4c0t2YjJiSW9lWDRYWmJzZmphZEY0R2hOUHdjalRROAE=",
      url: "paramountplus.com/movies/video/Y8sKvb2bIoeX4XZbsfjadF4GhNPwcjTQ",
   },
}

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
   for _, test := range tests {
      var pssh widevine.Pssh
      pssh.ContentId = []byte(test.content_id)
      pssh.KeyId, err = hex.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      var app AppToken
      app.ComCbsApp()
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
