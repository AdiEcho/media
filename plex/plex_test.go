package plex

import (
   "41.neocities.org/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestUrl(t *testing.T) {
   for _, test := range url_tests {
      var web Address
      web.Set(test)
      fmt.Println(web)
   }
}

var url_tests = []string{
   "/movie/the-hurt-locker",
   "/watch/movie/the-hurt-locker",
   "https://watch.plex.tv/watch/movie/the-hurt-locker",
   "watch.plex.tv/watch/movie/the-hurt-locker",
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
   var user Anonymous
   err = user.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range watch_tests {
      match, err := user.Match(&Address{test.path})
      if err != nil {
         t.Fatal(err)
      }
      video, err := user.Video(match, "")
      if err != nil {
         t.Fatal(err)
      }
      part, ok := video.Dash()
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
         var module widevine.Module
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
