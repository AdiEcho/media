package rtbf

import (
   "154.pages.dev/text"
   "41.neocities.org/widevine"
   "encoding/base64"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestAccountsLogin(t *testing.T) {
   username := os.Getenv("rtbf_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("rtbf_password")
   var login AuvioLogin
   err := login.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("login.txt", login.Raw, os.ModePerm)
}

func TestEntitlement(t *testing.T) {
   var (
      login AuvioLogin
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

func TestWebToken(t *testing.T) {
   var (
      login AuvioLogin
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
   token, err := login.Token()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", token)
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
