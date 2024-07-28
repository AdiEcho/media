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
   var login AuthLogin
   login.Data, err = os.ReadFile("login.json")
   if err != nil {
      t.Fatal(err)
   }
   login.Unmarshal()
   for _, film := range films {
      var pssh widevine.Pssh
      pssh.ContentId, err = base64.StdEncoding.DecodeString(film.content_id)
      if err != nil {
         t.Fatal(err)
      }
      pssh.KeyId, err = base64.StdEncoding.DecodeString(film.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Encode())
      if err != nil {
         t.Fatal(err)
      }
      movie, err := NewMovie(path.Base(film.url))
      if err != nil {
         t.Fatal(err)
      }
      title, err := login.Entitlement(movie)
      if err != nil {
         t.Fatal(err)
      }
      play, err := login.Playback(movie, title)
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(Poster{login, play}, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}

func TestLogin(t *testing.T) {
   username := os.Getenv("draken_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("draken_password")
   var login AuthLogin
   err := login.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("login.json", login.Data, 0666)
}
