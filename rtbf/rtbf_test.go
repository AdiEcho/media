package rtbf

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
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

var media = []struct{
   key_id string
   path   string
   url    string
}{
   {
      path: "/media/grantchester-grantchester-s01-3194636",
      url:  "auvio.rtbf.be/media/grantchester-grantchester-s01-3194636",
   },
   {
      path: "/emission/i-care-a-lot-27462",
      url:  "auvio.rtbf.be/emission/i-care-a-lot-27462",
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
   medium := media[0]
   var pssh widevine.Pssh
   pssh.KeyId, err = base64.StdEncoding.DecodeString(medium.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.Cdm
   err = module.New(private_key, client_id, pssh.Marshal())
   if err != nil {
      t.Fatal(err)
   }
   var login AuvioLogin
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
   var page AuvioPage
   err = page.New(medium.path)
   if err != nil {
      t.Fatal(err)
   }
   title, err := auth.Entitlement(&page)
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(title, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
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
   for _, medium := range media {
      var page AuvioPage
      err := page.New(medium.path)
      if err != nil {
         t.Fatal(err)
      }
      title, err := auth.Entitlement(&page)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", title)
      fmt.Println(title.Dash())
      time.Sleep(time.Second)
   }
}

func TestPage(t *testing.T) {
   for _, medium := range media {
      var page AuvioPage
      err := page.New(medium.path)
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
