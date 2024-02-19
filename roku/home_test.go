package roku

import (
   "154.pages.dev/encoding"
   "fmt"
   "testing"
   "time"
)

func TestContent(t *testing.T) {
   for _, test := range tests {
      var home HomeScreen
      err := home.New(test.playback_id)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(encoding.Name(home))
      time.Sleep(time.Second)
   }
}

const (
   episode = iota
   movie
)

var tests = map[int]struct {
   key string
   playback_id string
   pssh string
} {
   // therokuchannel.roku.com/watch/105c41ea75775968b670fbb26978ed76
   episode: {
      key: "e258b67d75420066c8424bd142f84565",
      playback_id: "105c41ea75775968b670fbb26978ed76",
      pssh: "AAAAQ3Bzc2gAAAAA7e+LqXnWSs6jyCfc1R0h7QAAACMIARIQvfpNbNs5cC5baB+QYX+afhoKaW50ZXJ0cnVzdCIBKg==",
   },
   // therokuchannel.roku.com/watch/597a64a4a25c5bf6af4a8c7053049a6f
   movie: {
      key: "13d7c7cf295444944b627ef0ad2c1b3c",
      playback_id: "597a64a4a25c5bf6af4a8c7053049a6f",
   },
}
