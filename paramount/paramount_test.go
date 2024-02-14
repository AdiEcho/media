package paramount

import (
   "fmt"
   "net/http"
   "path"
   "testing"
   "time"
)

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

func TestSecrets(t *testing.T) {
   test := tests["episode cenc"]
   for _, secret := range app_secrets {
      var at AppToken
      err := at.with(secret)
      if err != nil {
         t.Fatal(err)
      }
      if _, err := at.Item(path.Base(test.url)); err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
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
