package max

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestLogin(t *testing.T) {
   var login DefaultLogin
   login.Credentials.Username = os.Getenv("max_username")
   if login.Credentials.Username == "" {
      t.Fatal("Getenv")
   }
   login.Credentials.Password = os.Getenv("max_password")
   var key PublicKey
   err := key.New()
   if err != nil {
      t.Fatal(err)
   }
   var token DefaultToken
   err = token.New()
   if err != nil {
      t.Fatal(err)
   }
   err = token.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   err = token.Login(key, login)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("session.txt", token.Session.Raw, os.ModePerm)
   os.WriteFile("token.txt", token.Token.Raw, os.ModePerm)
}

func TestConfig(t *testing.T) {
   var token DefaultToken
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   decision, err := token.decision()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", decision)
}

func TestLicense(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   private_key, err := os.ReadFile(home + "/widevine/private_key.pem")
   if err != nil {
      t.Fatal(err)
   }
   client_id, err := os.ReadFile(home + "/widevine/client_id.bin")
   if err != nil {
      t.Fatal(err)
   }
   var token DefaultToken
   token.Session.Raw, err = os.ReadFile("session.txt")
   if err != nil {
      t.Fatal(err)
   }
   token.Token.Raw, err = os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = token.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var pssh widevine.Pssh
      pssh.KeyId, err = hex.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      var web Address
      web.UnmarshalText([]byte(test.url))
      play, err := token.Playback(web)
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(play, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}

var tests = []struct {
   key_id string
   url        string
   video_type string
}{
   {
      key_id: "0102f83e95fac43cf1662dd8e5b08d90",
      url: "play.max.com/video/watch/c9e9bde1-1463-4c92-a25a-21451f3c5894/f1d899ac-1780-494a-a20d-caee55c9e262",
      video_type: "MOVIE",
   },
   {
      key_id: "0102d949c44f81b28fdb98d535c8bade",
      url:        "play.max.com/video/watch/d0938760-d3ca-4c59-aea2-74ecbed42d17/2e7d1db4-2fd7-47fb-a7c3-a65b7c2e5d6f",
      video_type: "EPISODE",
   },
}

func TestRoutes(t *testing.T) {
   var token DefaultToken
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var web Address
      web.UnmarshalText([]byte(test.url))
      routes, err := token.Routes(web)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", routes)
      name, err := text.Name(routes)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
func TestLogin(t *testing.T) {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token bolt_token
   token.st = string(data)
   login, err := token.login()
   if err != nil {
      t.Fatal(err)
   }
   err = login.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", login.Data)
}
