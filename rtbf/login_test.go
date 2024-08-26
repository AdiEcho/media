package rtbf

import (
   "os"
   "testing"
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

var media = []struct {
   id     int64
   key_id string
   path   string
   url    string
}{
   {
      id:     3201987,
      key_id: "o1C37Tt5SzmHMmEgQViUEA==",
      path:   "/media/i-care-a-lot-i-care-a-lot-3201987",
      url:    "auvio.rtbf.be/media/i-care-a-lot-i-care-a-lot-3201987",
   },
   {
      path: "/media/grantchester-grantchester-s01-3194636",
      url:  "auvio.rtbf.be/media/grantchester-grantchester-s01-3194636",
   },
   {
      path: "/emission/i-care-a-lot-27462",
      url:  "auvio.rtbf.be/emission/i-care-a-lot-27462",
   },
}
