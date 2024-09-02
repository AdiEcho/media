package roku

import (
   "fmt"
   "os"
   "testing"
)

var tests = map[string]struct {
   id string
   key_id string
   url string
} {
   "episode": {
      id: "105c41ea75775968b670fbb26978ed76",
      key_id: "bdfa4d6cdb39702e5b681f90617f9a7e",
      url: "therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76",
   },
   "movie": {
      id: "597a64a4a25c5bf6af4a8c7053049a6f",
      key_id: "28339ad78f734520da24e6e0573d392e",
      url: "therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f",
   },
}

func TestTokenWrite(t *testing.T) {
   var err error
   // AccountAuth
   var auth AccountAuth
   auth.Raw, err = os.ReadFile("auth.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = auth.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   // AccountCode
   var code AccountCode
   code.Raw, err = os.ReadFile("code.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = code.Unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   // AccountToken
   token, err := auth.Token(&code)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.txt", token.Raw, os.ModePerm)
}

func TestTokenRead(t *testing.T) {
   var err      error
   // AccountToken
   var token AccountToken
   token.Raw, err = os.ReadFile("token.txt")
   if err != nil {
      t.Fatal(err)
   }
   token.Unmarshal()
   // AccountAuth
   var auth AccountAuth
   err = auth.New(&token)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", auth)
}
