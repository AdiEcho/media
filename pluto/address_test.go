package pluto

import (
   "154.pages.dev/text"
   "fmt"
   "testing"
   "time"
)

func TestAddress(t *testing.T) {
   for _, test := range video_tests {
      var web Address
      err := web.Set(test.url)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(web)
   }
}

func TestVideo(t *testing.T) {
   for _, test := range video_tests {
      var web Address
      web.Set(test.url)
      video, err := web.Video("")
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n", video)
      name, err := text.Name(Namer{video})
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(time.Second)
   }
}
