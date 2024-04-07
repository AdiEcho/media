package plex

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
)
package plex

import (
   "os"
   "testing"
   "time"
)

var paths = []string{
   //"/movie/cruel-intentions",
   "/show/broadchurch/season/3/episode/5",
}

func TestMatch(t *testing.T) {
   var anon anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, path := range paths {
      func() {
         res, err := anon.matches(path)
         if err != nil {
            t.Fatal(err)
         }
         defer res.Body.Close()
         res.Write(os.Stdout)
      }()
      time.Sleep(time.Second)
   }
}
const (
   cruel_intentions = "https://watch.plex.tv/movie/cruel-intentions"
   default_kid = "eabdd790d9279b9699b32110eed9a154"
)

func TestMetadata(t *testing.T) {
   var anon anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   var web web_address
   if err := web.Set(cruel_intentions); err != nil {
      t.Fatal(err)
   }
   meta, err := anon.metadata(web)
   if err != nil {
      t.Fatal(err)
   }
   part, ok := meta.dash(anon)
   if !ok {
      t.Fatal("metadata.dash")
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
   key_id, err := hex.DecodeString(default_kid)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   if err := module.New(private_key, client_id, key_id); err != nil {
      t.Fatal(err)
   }
   license, err := module.License(part)
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Printf("%x %v\n", key, ok)
}
