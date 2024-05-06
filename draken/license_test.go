package draken

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "path"
   "testing"
   "time"
)

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
   var auth auth_login
   auth.data, err = os.ReadFile("login.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.unmarshal()
   for _, film := range films {
      key_id, err := base64.StdEncoding.DecodeString(film.key_id)
      if err != nil {
         t.Fatal(err)
      }
      content_id, err := base64.StdEncoding.DecodeString(film.content_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.CDM
      err = module.New(private_key, client_id, widevine.PSSH(key_id, content_id))
      if err != nil {
         t.Fatal(err)
      }
      movie, err := new_movie(path.Base(film.url))
      if err != nil {
         t.Fatal(err)
      }
      title, err := auth.entitlement(movie)
      if err != nil {
         t.Fatal(err)
      }
      play, err := auth.playback(movie, title)
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(poster{auth, play}, key_id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}
