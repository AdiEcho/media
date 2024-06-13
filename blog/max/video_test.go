package max

import (
   "fmt"
   "testing"
)

func TestVideo(t *testing.T) {
   var token default_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   video, err := token.video()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
}
