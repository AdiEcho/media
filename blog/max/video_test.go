package max

import (
   "154.pages.dev/text"
   "fmt"
   "testing"
)

// play.max.com/movie/127b00c5-0131-4bac-b2d1-40762deefe09
const show = "127b00c5-0131-4bac-b2d1-40762deefe09"

func TestVideo(t *testing.T) {
   var token default_token
   err := token.New()
   if err != nil {
      t.Fatal(err)
   }
   video, err := token.video(show)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
   name, err := text.Name(video)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", name)
}
