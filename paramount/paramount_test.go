package paramount

import (
   "154.pages.dev/widevine"
   "encoding/base64"
   "fmt"
   "net/http"
   "os"
   "path"
   "testing"
   "time"
)

func TestPost(t *testing.T) {
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
   test := tests["episode cenc"]
   pssh, err := base64.StdEncoding.DecodeString(test.pssh)
   if err != nil {
      t.Fatal(err)
   }
   mod, err := widevine.NewModule(private_key, client_id, nil, pssh)
   if err != nil {
      t.Fatal(err)
   }
   token, err := NewAppToken()
   if err != nil {
      t.Fatal(err)
   }
   sess, err := token.Session(path.Base(test.url))
   if err != nil {
      t.Fatal(err)
   }
   key, err := mod.Key(sess)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%x\n", key)
}

func TestSecrets(t *testing.T) {
   test := tests["episode cenc"]
   for _, secret := range app_secrets {
      token, err := app_token_with(secret)
      if err != nil {
         t.Fatal(err)
      }
      if _, err := token.Item(path.Base(test.url)); err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}

var tests = map[string]struct{
   asset func(string)(string,error)
   pssh string
   url string
}{
   "episode cenc": {
      asset: DashCenc,
      pssh: "AAAAWHBzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAADgIARIQD3gqa9LyRm65UzN2yiD/XyIgcm4xenlpclZPUGpDbDhyeG9wV3JoVW1KRUlzM0djS1c4AQ==",
      url: "paramountplus.com/shows/video/rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW",
   },
   "episode downloadable": {
      asset: Downloadable,
      url: "paramountplus.com/shows/video/YxlqOUdP1zZaIs7FGXCaS1dJi7gGzxG_",
   },
   "movie cenc": {
      asset: DashCenc,
      url: "paramountplus.com/movies/video/tQk_Qooh5wUlxQqzj_4LiBO2m4iMrcPD",
   },
}

func TestMedia(t *testing.T) {
   for _, test := range tests {
      ref, err := test.asset(path.Base(test.url))
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(ref)
      func() {
         res, err := http.Get(ref)
         if err != nil {
            t.Fatal(err)
         }
         defer res.Body.Close()
         if res.StatusCode != http.StatusOK {
            if res.StatusCode != http.StatusFound {
               t.Fatal(res)
            }
         }
      }()
      time.Sleep(time.Second)
   }
}
