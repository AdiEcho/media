package tubi

import (
   "154.pages.dev/widevine"
   "encoding/hex"
   "fmt"
   "os"
   "testing"
   "time"
)

var tests = map[string]struct{
   content_id int
   key_id string
   url string
}{
   "movie": {
      content_id: 589292,
      key_id: "943974887f2a4b87a3ded9e99f03c962",
      url: "tubitv.com/movies/589292",
   },
   "episode": {
      content_id: 200042567,
      url: "tubitv.com/tv-shows/200042567",
   },
   "series": {
      content_id: 300002169,
      url: "tubitv.com/series/300002169",
   },
}

func TestContent(t *testing.T) {
   for _, test := range tests {
      var content content_management
      err := content.New(test.content_id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", content)
      time.Sleep(time.Second)
   }
}

func TestLicense(t *testing.T) {
   test := tests["movie"]
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
   key_id, err := hex.DecodeString(test.key_id)
   if err != nil {
      t.Fatal(err)
   }
   var module widevine.CDM
   if err := module.New(private_key, client_id, key_id); err != nil {
      t.Fatal(err)
   }
   var content content_management
   if err := content.New(test.content_id); err != nil {
      t.Fatal(err)
   }
   video, ok := content.Resolution720p()
   if !ok {
      t.Fatal("Resolution720p")
   }
   license, err := module.License(video)
   if err != nil {
      t.Fatal(err)
   }
   key, ok := module.Key(license)
   fmt.Printf("%x %v\n", key, ok)
}
