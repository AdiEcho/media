package ctv

import (
   "fmt"
   "os"
   "testing"
)

// ctv.ca/movies/the-girl-with-the-dragon-tattoo-2011
const dragon_tattoo = "/movies/the-girl-with-the-dragon-tattoo-2011"

func TestWrite(t *testing.T) {
   var path resolve_path
   err := path.New(dragon_tattoo)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("resolvePath.json", path.data, 0666)
   err = path.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   packages, err := path.v.Data.ResolvedPath.LastSegment.packages()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("contentPackages.json", packages.data, 0666)
}

func TestRead(t *testing.T) {
   var (
      path resolve_path
      err error
   )
   path.data, err = os.ReadFile("resolvePath.json")
   if err != nil {
      t.Fatal(err)
   }
   path.unmarshal()
   var packages content_packages
   packages.data, err = os.ReadFile("contentPackages.json")
   if err != nil {
      t.Fatal(err)
   }
   packages.unmarshal()
   manifest := path.v.Data.ResolvedPath.LastSegment.manifest(packages)
   fmt.Printf("%q\n", manifest)
}
