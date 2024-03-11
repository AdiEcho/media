package web

import (
   "154.pages.dev/media/blog/spotify/android"
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
)

const pssh = "AAAAU3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADMIARIQOSSC/pvtc3LRZX1+IvMreRoHc3BvdGlmeSIUOSSC/pvtc3LRZX1+IvMreSkC870="

func TestLicense(t *testing.T) {
   data, err := base64.StdEncoding.DecodeString(pssh)
   if err != nil {
      t.Fatal(err)
   }
   var protect widevine.PSSH
   if err := protect.New(data); err != nil {
      t.Fatal(err)
   }
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
   module, err := protect.CDM(private_key, client_id)
   if err != nil {
      t.Fatal(err)
   }
   var login android.LoginOk
   login.Data, err = os.ReadFile(home + "/spotify.bin")
   if err != nil {
      t.Fatal(err)
   }
   if err := login.Consume(); err != nil {
      t.Fatal(err)
   }
   license, err := module.License(login)
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Printf("%x %v\n", key, ok)
}

func TestSeektable(t *testing.T) {
   var seek seektable
   err := seek.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", seek)
}
