package spotify

import (
   "fmt"
   "os"
   "testing"
)

func TestPhp(t *testing.T) {
   data, err := os.ReadFile(".txt")
   if err != nil {
      t.Fatal(err)
   }
   var login login_response
   _, data, _ = bytes.Cut(data, []byte("\r\n\r\n"))
   if err := login.m.Consume(data); err != nil {
      t.Fatal(err)
   }
   message, err := login.challenge(username, password)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%#v\n", message)
}

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
