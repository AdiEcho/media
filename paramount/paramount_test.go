package paramount

import (
   "fmt"
   "net/http"
   "path"
   "testing"
   "time"
)

func TestSecrets(t *testing.T) {
   test := tests["episode"]
   for _, secret := range app_secrets {
      var at AppToken
      err := at.with(secret)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(secret)
      if _, err := at.Item(path.Base(test.url)); err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
   }
}

var tests = map[string]struct{
   asset func(string)(string,error)
   key_id string
   url string
}{
   "episode": {
      asset: DashCenc,
      key_id: "0f782a6bd2f2466eb9533376ca20ff5f",
      url: "paramountplus.com/shows/video/rn1zyirVOPjCl8rxopWrhUmJEIs3GcKW",
   },
   "movie": {
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
