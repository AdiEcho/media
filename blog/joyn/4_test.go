package joyn

import (
   "os"
   "testing"
)

func TestPlaylist(t *testing.T) {
   var anon anonymous
   err := anon.New()
   if err != nil {
      t.Fatal(err)
   }
   var movie movie_detail
   err = movie.New(barry_seal)
   if err != nil {
      t.Fatal(err)
   }
   title, err := anon.entitlement(movie)
   if err != nil {
      t.Fatal(err)
   }
   res, err := title.playlist(movie)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
