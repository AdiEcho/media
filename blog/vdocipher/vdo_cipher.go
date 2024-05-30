package main

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "encoding/json"
   "fmt"
   "net/http"
   "os"
)

type poster struct{}

func (poster) RequestUrl() (string, bool) {
   return "https://license.vdocipher.com/auth", true
}

func (poster) RequestHeader() (http.Header, error) {
   h := make(http.Header)
   h.Set("content-type", "application/json")
   return h, nil
}

func (poster) WrapRequest(b []byte) ([]byte, error) {
   inner, err := func() ([]byte, error) {
      var s struct {
         LicenseRequest []byte `json:"licenseRequest"`
         OTP string `json:"otp"`
         Tech string `json:"tech"`
      }
      s.LicenseRequest = b
      s.OTP = "20160313versASE323yAKNwZ7gSdMScMmqg2jfdmHByn89Koc6N1jMSjXcBkMFO1"
      s.Tech = "wv"
      return json.Marshal(s)
   }()
   if err != nil {
      return nil, err
   }
   var s struct {
      Token []byte `json:"token"`
   }
   s.Token = inner
   return json.Marshal(s)
}

func (poster) UnwrapResponse(b []byte) ([]byte, error) {
   var s struct {
      License []byte
   }
   err := json.Unmarshal(b, &s)
   if err != nil {
      return nil, err
   }
   return s.License, nil
}

func main() {
   home, err := os.UserHomeDir()
   if err != nil {
      panic(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      panic(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      panic(err)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, widevine.PSSH(
      nil, []byte("vdocipher:86bcd779dbbf49c9a9c6c0ecc19ffb7d"),
   ))
   if err != nil {
      panic(err)
   }
   key_id, err := hex.DecodeString("3cc487be2e565228ab77a8677e29943d")
   if err != nil {
      panic(err)
   }
   key, err := module.Key(poster{}, key_id)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%x\n", key)
}
