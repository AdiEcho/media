package ctv

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

// ctv.ca/movies/the-girl-with-the-dragon-tattoo-2011
const (
   raw_content_id = "ZmYtZDAxM2NhN2EtMjY0MjY1"
   raw_key_id = "ywlXHuvLP3KHICZX9rn3pg=="
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
   content_id, err := base64.StdEncoding.DecodeString(raw_content_id)
   if err != nil {
      t.Fatal(err)
   }
   key_id, err := base64.StdEncoding.DecodeString(raw_key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, widevine.PSSH(
      nil, content_id,
   ))
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(Poster{}, key_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

var test_paths = []string{
   // ctv.ca/shows/friends/the-one-with-the-chicken-pox-s2e23
   "/shows/friends/the-one-with-the-chicken-pox-s2e23",
   // ctv.ca/movies/the-girl-with-the-dragon-tattoo-2011
   "/movies/the-girl-with-the-dragon-tattoo-2011",
   // ctv.ca/movies/baby-driver
   "/movies/baby-driver",
}

func TestMedia(t *testing.T) {
   for _, test_path := range test_paths {
      resolve, err := Path(test_path).Resolve()
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
      axis, err := resolve.Axis()
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
      media, err := axis.Media()
      if err != nil {
         t.Fatal(err)
      }
      name, err := text.Name(MediaManifest{M: media})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(99 * time.Millisecond)
   }
}

func TestManifest(t *testing.T) {
   for _, test_path := range test_paths {
      resolve, err := Path(test_path).Resolve()
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
      axis, err := resolve.Axis()
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
      media, err := axis.Media()
      if err != nil {
         t.Fatal(err)
      }
      manifest, err := axis.Manifest(media)
      if err != nil {
         t.Fatal(err)
      }
      text, err := manifest.Marshal()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(string(text))
      time.Sleep(99 * time.Millisecond)
   }
}
