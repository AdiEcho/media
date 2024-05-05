package draken

import (
   "bytes"
   "fmt"
   "os"
   "testing"
)

var custom_ids = []string{
   // drakenfilm.se/film/michael-clayton
   "michael-clayton",
   // drakenfilm.se/film/the-card-counter
   "the-card-counter",
}

func TestLicense(t *testing.T) {
   var (
      auth auth_login
      err error
   )
   auth.data, err = os.ReadFile("login.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.unmarshal()
   movie, err := new_movie(custom_ids[0])
   if err != nil {
      t.Fatal(err)
   }
   title, err := auth.entitlement(movie)
   if err != nil {
      t.Fatal(err)
   }
   play, err := auth.playback(movie, title)
   if err != nil {
      t.Fatal(err)
   }
   res, err := auth.license(play)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   buf := new(bytes.Buffer)
   res.Write(buf)
   fmt.Printf("%q\n", buf)
}
