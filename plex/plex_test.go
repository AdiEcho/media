package plex

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestVideo(t *testing.T) {
   var anon Anonymous
   err = anon.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range watch_tests {
      match, err := anon.Match(Url{test.path})
      if err != nil {
         t.Fatal(err)
      }
      video, err := anon.Video(match, "")
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", video)
      time.Sleep(time.Second)
   }
}

func TestUrl(t *testing.T) {
   for _, test := range url_tests {
      var u Url
      u.Set(test)
      fmt.Println(u)
   }
}

var url_tests = []string{
   "/movie/the-hurt-locker",
   "/watch/movie/the-hurt-locker",
   "https://watch.plex.tv/watch/movie/the-hurt-locker",
   "watch.plex.tv/watch/movie/the-hurt-locker",
}

var watch_tests = []struct{
   key_id string
   path string
   url string
}{
   {
      key_id: "4310a7c8094acab73fceab9d5494f36f",
      path: "/movie/cruel-intentions",
      url: "watch.plex.tv/movie/cruel-intentions",
   },
   {
      key_id: "", // no DRM
      path: "/show/broadchurch/season/3/episode/5",
      url: "watch.plex.tv/show/broadchurch/season/3/episode/5",
   },
}

func TestDiscover(t *testing.T) {
   var anon Anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range watch_tests {
      match, err := anon.Match(Url{test.path})
      if err != nil {
         t.Fatal(err)
      }
      name, err := text.Name(&Namer{match})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(name)
      time.Sleep(time.Second)
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
   var anon Anonymous
   err = anon.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range watch_tests {
      match, err := anon.Match(Url{test.path})
      if err != nil {
         t.Fatal(err)
      }
      video, err := anon.Video(match, "")
      if err != nil {
         t.Fatal(err)
      }
      part, ok := video.Dash(anon)
      if !ok {
         t.Fatal("Metadata.Dash")
      }
      fmt.Printf("%+v\n", part)
      if test.key_id != "" {
         var pssh widevine.Pssh
         pssh.KeyId, err = hex.DecodeString(test.key_id)
         if err != nil {
            t.Fatal(err)
         }
         var module widevine.Cdm
         err = module.New(private_key, client_id, pssh.Marshal())
         if err != nil {
            t.Fatal(err)
         }
         key, err := module.Key(part, pssh.KeyId)
         if err != nil {
            t.Fatal(err)
         }
         fmt.Printf("%x\n", key)
      }
      time.Sleep(time.Second)
   }
}
