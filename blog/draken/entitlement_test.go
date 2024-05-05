package draken

import (
   "fmt"
   "os"
   "testing"
   "time"
)

func TestEntitlement(t *testing.T) {
   username := os.Getenv("draken_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("draken_password")
   var auth auth_login
   err := auth.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   for _, id := range custom_ids {
      movie, err := new_movie(id)
      if err != nil {
         t.Fatal(err)
      }
      title, err := auth.entitlement(movie)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", title)
      time.Sleep(time.Second)
   }
}
