package spotify

import (
   "fmt"
   "os"
   "testing"
)

func TestOkWrite(t *testing.T) {
   username := os.Getenv("spotify_username")
   if username == "" {
      t.Fatal("spotify_username")
   }
   password := os.Getenv("spotify_password")
   var response login_response
   err := response.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   ok, err := response.ok(username, password)
   if err != nil {
      t.Fatal(err)
   }
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile(home + "/spotify.bin", ok.data, 0666)
}

func TestOkRead(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var login login_ok
   login.data, err = os.ReadFile(home + "/spotify.bin")
   if err != nil {
      t.Fatal(err)
   }
   if err := login.consume(); err != nil {
      t.Fatal(err)
   }
   fmt.Println(login.access_token())
}
