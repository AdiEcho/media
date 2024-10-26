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

func TestLogin(t *testing.T) {
   data, err := os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   var token BoltToken
   token.St = string(data)
   login, err := token.Login()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("login.txt", login.Raw, os.ModePerm)
}

func TestRoutes(t *testing.T) {
   var (
      login LinkLogin
      err error
   )
   login.Raw, err = os.ReadFile("login.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = login.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var web Address
      web.UnmarshalText([]byte(test.url))
      routes, err := login.Routes(web)
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
   var login LinkLogin
   login.Raw, err = os.ReadFile("login.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = login.Unmarshal()
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
      play, err := login.Playback(web)
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
