package max

import (
   "fmt"
   "os"
   "testing"
)

func TestAndroidPlayback(t *testing.T) {
   text, err := os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   var token default_token
   err = token.unmarshal(text)
   if err != nil {
      t.Fatal(err)
   }
   var playback playback_request
   playback.New()
   resp, err := token.playback(playback)
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

func TestAndroidLogin(t *testing.T) {
   var login default_login
   login.Credentials.Username = os.Getenv("max_username")
   if login.Credentials.Username == "" {
      t.Fatal("Getenv")
   }
   login.Credentials.Password = os.Getenv("max_password")
   var key public_key
   err := key.New()
   if err != nil {
      t.Fatal(err)
   }
   var token default_token
   err = token.New()
   if err != nil {
      t.Fatal(err)
   }
   err = token.login(key, login)
   if err != nil {
      t.Fatal(err)
   }
   text, err := token.marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.json", text, 0666)
}

func TestAndroidConfig(t *testing.T) {
   var token default_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   config, err := token.config()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%s\n", config)
}
