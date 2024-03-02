package spotify

import (
   "fmt"
   "os"
   "testing"
)

func TestChallenge(t *testing.T) {
   username := os.Getenv("spotify_username")
   password := os.Getenv("spotify_password")
   var login login_response
   err := login.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   message, err := login.challenge(username, password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%#v\n", message)
}
