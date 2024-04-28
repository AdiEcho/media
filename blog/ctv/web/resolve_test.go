package ctv

import (
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
