package ctv

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
)

// ctv.ca/shows/friends/the-one-with-the-bullies-s2e21
func TestWriteEpisode(t *testing.T) {
   err := write("/shows/friends/the-one-with-the-bullies-s2e21")
   if err != nil {
      t.Fatal(err)
   }
}

// ctv.ca/movies/the-girl-with-the-dragon-tattoo-2011
func TestWriteMovie(t *testing.T) {
   err := write("/movies/the-girl-with-the-dragon-tattoo-2011")
   if err != nil {
      t.Fatal(err)
   }
}

func TestRead(t *testing.T) {
   manifest, err := read()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", manifest)
}

func read() (string, error) {
   var (
      resolve resolve_path
      err error
   )
   resolve.data, err = os.ReadFile("resolvePath.json")
   if err != nil {
      return "", err
   }
   resolve.unmarshal()
   var packages content_packages
   packages.data, err = os.ReadFile("contentPackages.json")
   if err != nil {
      return "", err
   }
   packages.unmarshal()
   return resolve.v.Data.ResolvedPath.LastSegment.manifest(packages), nil
}

func write(path string) error {
   var resolve resolve_path
   err := resolve.New(path)
   if err != nil {
      return err
   }
   os.WriteFile("resolvePath.json", resolve.data, 0666)
   err = resolve.unmarshal()
   if err != nil {
      return err
   }
   packages, err := resolve.v.Data.ResolvedPath.LastSegment.packages()
   if err != nil {
      return err
   }
   return os.WriteFile("contentPackages.json", packages.data, 0666)
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
   license, err := module.License(poster{})
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(license)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}
