package rtbf

import (
   "154.pages.dev/text"
   "fmt"
   "os"
   "testing"
   "time"
)

var media = []struct{
   id int64
   key_id string
   path string
   url string
}{
   {
      id: 3201987,
      key_id: "o1C37Tt5SzmHMmEgQViUEA==",
      path: "/media/i-care-a-lot-i-care-a-lot-3201987",
      url: "auvio.rtbf.be/media/i-care-a-lot-i-care-a-lot-3201987",
   },
   {
      path: "/emission/i-care-a-lot-27462",
      url: "auvio.rtbf.be/emission/i-care-a-lot-27462",
   },
   {
      path: "/media/grantchester-grantchester-s01-3194636",
      url: "auvio.rtbf.be/media/grantchester-grantchester-s01-3194636",
   },
}

func TestEmbedMedia(t *testing.T) {
   for _, medium := range media {
      var embed embed_media
      err := embed.New(medium.id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", embed)
      name, err := text.Name(embed)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}

func TestWebToken(t *testing.T) {
   text, err := os.ReadFile("account.json")
   if err != nil {
      t.Fatal(err)
   }
   var account accounts_login
   err = account.unmarshal(text)
   if err != nil {
      t.Fatal(err)
   }
   token, err := account.token()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", token)
}

func TestAccountsLogin(t *testing.T) {
   username, password := os.Getenv("rtbf_username"), os.Getenv("rtbf_password")
   if username == "" {
      t.Fatal("Getenv")
   }
   var login accounts_login
   err := login.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   text, err := login.marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("account.json", text, 0666)
}