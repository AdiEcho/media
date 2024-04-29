package ctv

import (
   "154.pages.dev/encoding"
   "encoding/json"
   "fmt"
   "os"
   "testing"
   "time"
)

var test_paths = []string{
   // ctv.ca/movies/the-girl-with-the-dragon-tattoo-2011
   "/movies/the-girl-with-the-dragon-tattoo-2011",
   // ctv.ca/shows/friends/the-one-with-the-bullies-s2e21
   "/shows/friends/the-one-with-the-bullies-s2e21",
}

func TestMedia(t *testing.T) {
   enc := json.NewEncoder(os.Stdout)
   enc.SetIndent("", " ")
   for _, path := range test_paths {
      resolve, err := new_resolve(path)
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
      content, err := resolve.content()
      if err != nil {
         t.Fatal(err)
      }
      time.Sleep(99 * time.Millisecond)
      media, err := content.media()
      if err != nil {
         t.Fatal(err)
      }
      name, err := encoding.Name(media)
      if err != nil {
         t.Fatal(err)
      }
      err = enc.Encode(media)
      if err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%q\n", name)
      time.Sleep(99 * time.Millisecond)
   }
}
