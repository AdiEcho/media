package mubi

import (
   "os"
   "testing"
)

// mubi.com/films/dogville
const dogville = "/films/dogville"

func TestFilms(t *testing.T) {
   res, err := films(dogville)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
