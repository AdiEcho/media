package draken

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "path"
   "testing"
   "time"
)

var films = []struct {
   content_id string
   key_id     string
   url        string
}{
   {
      content_id: "ODE0OTQ1NWMtY2IzZC00YjE1LTg1YTgtYjk1ZTNkMTU3MGI1",
      key_id:     "e5WypDjIM1+4W74cf6rHIw==",
      url:        "drakenfilm.se/film/michael-clayton",
   },
   {
      content_id: "MTcxMzkzNTctZWQwYi00YTE2LThiZTYtNjllNDE4YzRiYTQw",
      key_id:     "ToV4wH2nlVZE8QYLmLywDg==",
      url:        "drakenfilm.se/film/the-card-counter",
   },
}

func TestMovie(t *testing.T) {
   for _, film := range films {
      movie, err := NewMovie(path.Base(film.url))
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", movie)
      name, err := text.Name(Namer{movie})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(99 * time.Millisecond)
   }
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
