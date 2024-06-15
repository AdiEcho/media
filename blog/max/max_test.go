package max

import (
   "154.pages.dev/text"
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = []struct{
   video_type string
   path string
   url string
}{
   {
      video_type: "MOVIE",
      path: "/video/watch/b3b1410a-0c85-457b-bcc7-e13299bea2a8/1623fe4c-ef6e-4dd1-a10c-4a181f5f6579",
      url: "play.max.com/video/watch/b3b1410a-0c85-457b-bcc7-e13299bea2a8/1623fe4c-ef6e-4dd1-a10c-4a181f5f6579",
   },
   {
      video_type: "EPISODE",
      path: "/video/watch/fbdd33a2-1189-4b9a-8c10-13244fb21b7f/6cc15a42-130f-4531-807a-b2c147d8ac68",
      url: "play.max.com/video/watch/fbdd33a2-1189-4b9a-8c10-13244fb21b7f/6cc15a42-130f-4531-807a-b2c147d8ac68",
   },
}

func TestRoutes(t *testing.T) {
   var token default_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      routes, err := token.routes(test.path)
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
   var pssh widevine.PSSH
   pssh.KeyId, err = hex.DecodeString(default_kid)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   err = module.New(private_key, client_id, pssh.Encode())
   if err != nil {
      t.Fatal(err)
   }
   text, err := os.ReadFile("token.json")
   if err != nil {
      t.Fatal(err)
   }
   var token default_token
   token.unmarshal(text)
   var web address
   web.UnmarshalText([]byte(tests[0].path))
   play, err := token.playback(web)
   if err != nil {
      t.Fatal(err)
   }
   key, err := module.Key(play, pssh.KeyId)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

const default_kid = "01021e5f16aa2c5ed02c550139b5ab82"
