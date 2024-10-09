package draken

import (
   "154.pages.dev/text"
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "reflect"
   "testing"
   "time"
)

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   AuthLogin{},
   Entitlement{},
   FullMovie{},
   Namer{},
   Playback{},
   Poster{},
   header{},
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
   os.WriteFile("login.txt", login.Raw, os.ModePerm)
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
   login.Raw, err = os.ReadFile("login.txt")
   if err != nil {
      t.Fatal(err)
   }
   if err = login.Unmarshal(); err != nil {
      t.Fatal(err)
   }
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
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      var movie FullMovie
      err = movie.New(film.custom_id)
      if err != nil {
         t.Fatal(err)
      }
      title, err := login.Entitlement(&movie)
      if err != nil {
         t.Fatal(err)
      }
      play, err := login.Playback(&movie, title)
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(&Poster{&login, play}, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}

func TestMovie(t *testing.T) {
   for _, film := range films {
      var movie FullMovie
      if err := movie.New(film.custom_id); err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", movie)
      name, err := text.Name(&Namer{&movie})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(99 * time.Millisecond)
   }
}

var films = []struct {
   content_id string
   custom_id string
   key_id     string
   url        string
}{
   {
      content_id: "ODE0OTQ1NWMtY2IzZC00YjE1LTg1YTgtYjk1ZTNkMTU3MGI1",
      custom_id: "michael-clayton",
      key_id:     "e5WypDjIM1+4W74cf6rHIw==",
      url:        "drakenfilm.se/film/michael-clayton",
   },
   {
      content_id: "MTcxMzkzNTctZWQwYi00YTE2LThiZTYtNjllNDE4YzRiYTQw",
      custom_id:        "the-card-counter",
      key_id:     "ToV4wH2nlVZE8QYLmLywDg==",
      url:        "drakenfilm.se/film/the-card-counter",
   },
}
