package ctv

import (
   "154.pages.dev/encoding"
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestManifest(t *testing.T) {
   for _, test_path := range test_paths {
      resolve, err := NewResolve(test_path)
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
      text, err := manifest.marshal()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(string(text))
      time.Sleep(99 * time.Millisecond)
   }
}

// ctv.ca/movies/the-girl-with-the-dragon-tattoo-2011
const raw_pssh = "CAESEMsJVx7ryz9yhyAmV/a596YaCWJlbGxtZWRpYSISZmYtZDAxM2NhN2EtMjY0MjY1"

func TestLicense(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   pssh, err := base64.StdEncoding.DecodeString(raw_pssh)
   if err != nil {
      t.Fatal(err)
   }
   module, err := widevine.PSSH(pssh).CDM(client_id, private_key)
   if err != nil {
      t.Fatal(err)
   }
   license, err := module.License(Poster{})
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(license)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
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
      name, err := encoding.Name(Namer{media})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(99 * time.Millisecond)
   }
}

var test_paths = []string{
   // ctv.ca/movies/the-girl-with-the-dragon-tattoo-2011
   "/movies/the-girl-with-the-dragon-tattoo-2011",
   // ctv.ca/shows/friends/the-one-with-the-bullies-s2e21
   "/shows/friends/the-one-with-the-bullies-s2e21",
}

