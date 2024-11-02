package rtbf

import (
   "41.neocities.org/text"
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "reflect"
   "strings"
   "testing"
   "time"
)

func TestWidevine(t *testing.T) {
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
   var login AuvioLogin
   login.Raw, err = os.ReadFile(home + "/rtbf.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = login.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   token, err := login.Token()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := token.Auth()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      var pssh widevine.Pssh
      pssh.KeyId, err = base64.StdEncoding.DecodeString(test.key_id)
      if err != nil {
         t.Fatal(err)
      }
      var module widevine.Cdm
      err = module.New(private_key, client_id, pssh.Marshal())
      if err != nil {
         t.Fatal(err)
      }
      page, err := Address{test.path}.Page()
      if err != nil {
         t.Fatal(err)
      }
      asset_id, ok := page.GetAssetId()
      if !ok {
         t.Fatal("AuvioPage.GetAssetId")
      }
      title, err := auth.Entitlement(asset_id)
      if err != nil {
         t.Fatal(err)
      }
      key, err := module.Key(title, pssh.KeyId)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%x\n", key)
      time.Sleep(time.Second)
   }
}
func TestAccountsLogin(t *testing.T) {
   username, password, ok := strings.Cut(os.Getenv("rtbf"), ":")
   if !ok {
      t.Fatal("Getenv")
   }
   data, err := AuvioLogin{}.Marshal(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("login.txt", data, os.ModePerm)
}

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Address{},
   AuvioAuth{},
   AuvioLogin{},
   AuvioPage{},
   Entitlement{},
   Namer{},
   Subtitle{},
   Title{},
   WebToken{},
}

var tests = []struct{
   key_id string
   path   string
   url    string
}{
   {
      key_id: "v6f4GpHISrWZcd6esgIvRw==",
      path: "/emission/la-proposition-27866",
      url: "auvio.rtbf.be/emission/la-proposition-27866",
   },
   {
      key_id: "10kWa4A9SOSzHFpq1n1zUQ==",
      path: "/media/grantchester-grantchester-s01-3194636",
      url:  "auvio.rtbf.be/media/grantchester-grantchester-s01-3194636",
   },
}
func TestWebToken(t *testing.T) {
   data, err := os.ReadFile("login.txt")
   if err != nil {
      t.Fatal(err)
   }
   var login AuvioLogin
   err = login.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   token, err := login.Token()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", token)
}

func TestEntitlement(t *testing.T) {
   data, err := os.ReadFile("login.txt")
   if err != nil {
      t.Fatal(err)
   }
   var login AuvioLogin
   err = login.Unmarshal(data)
   if err != nil {
      t.Fatal(err)
   }
   token, err := login.Token()
   if err != nil {
      t.Fatal(err)
   }
   auth, err := token.Auth()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      page, err := Address{test.path}.Page()
      if err != nil {
         t.Fatal(err)
      }
      asset_id, ok := page.GetAssetId()
      if !ok {
         t.Fatal("AuvioPage.GetAssetId")
      }
      title, err := auth.Entitlement(asset_id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", title)
      fmt.Println(title.Dash())
      time.Sleep(time.Second)
   }
}

func TestPage(t *testing.T) {
   for _, test := range tests {
      page, err := Address{test.path}.Page()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", page)
      name, err := text.Name(&Namer{page})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
