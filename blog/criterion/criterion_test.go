package criterion

import (
   "154.pages.dev/encoding"
   "fmt"
   "os"
   "testing"
)

// criterionchannel.com/videos/my-dinner-with-andre
const video_id = 455774

func TestVideo(t *testing.T) {
   var (
      token auth_token
      err error
   )
   token.data, err = os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   token.unmarshal()
   video, err := token.video(video_id)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
   name, err := encoding.Name(namer{video})
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", name)
}

func TestToken(t *testing.T) {
   username := os.Getenv("criterion_username")
   if username == "" {
      t.Fatal("Getenv")
   }
   password := os.Getenv("criterion_password")
   var token auth_token
   err := token.New(username, password)
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("token.json", token.data, 0666)
}
