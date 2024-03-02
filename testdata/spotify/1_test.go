package spotify

import (
   "fmt"
   "os"
   "testing"
)

func TestLogin(t *testing.T) {
   username := os.Getenv("spotify_username")
   password := os.Getenv("spotify_password")
   var req login_request
   if !req.New(username, password) {
      t.Fatal("Getenv")
   }
   res, err := req.login()
   if err != nil {
      t.Fatal(err)
   }
   suffix, duration, err := res.solve_hash_cash_challenge()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(suffix, duration)
}
