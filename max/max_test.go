package max

import (
   "154.pages.dev/text"
   "encoding/json"
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
   err = token.Login(key, login)
   if err != nil {
      t.Fatal(err)
   }
   text, err := token.Marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.json", text, 0666)
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
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   enc.Encode(decision)
}

func TestToken(t *testing.T) {
   var token DefaultToken
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", token)
}

func TestRoutes(t *testing.T) {
   var token DefaultToken
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var web WebAddress
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
