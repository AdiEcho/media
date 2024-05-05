package draken

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestPlayback(t *testing.T) {
   var (
      auth auth_login
      err error
   )
   auth.data, err = os.ReadFile("login.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.unmarshal()
   for _, id := range custom_ids {
      movie, err := new_movie(id)
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
      fmt.Printf("%+v\n", play)
      time.Sleep(time.Second)
   }
}
