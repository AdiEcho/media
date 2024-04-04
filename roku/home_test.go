package roku

import (
   "154.pages.dev/encoding"
   "fmt"
   "path"
   "testing"
   "time"
)

var tests = map[string]struct {
   key string
   key_id string
   url string
} {
   "episode": {
      key: "e258b67d75420066c8424bd142f84565",
      key_id: "bdfa4d6cdb39702e5b681f90617f9a7e",
      url: "therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76",
   },
   "movie": {
      key: "13d7c7cf295444944b627ef0ad2c1b3c",
      url: "therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f",
   },
}

func TestContent(t *testing.T) {
   for _, test := range tests {
      var home HomeScreen
      err := home.New(path.Base(test.url))
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(encoding.Name(home))
      time.Sleep(time.Second)
   }
}
