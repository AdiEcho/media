package amc

import (
   "154.pages.dev/encoding"
   "fmt"
   "os"
   "testing"
   "time"
)

func TestContent(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var auth Authorization
   auth.Raw, err = os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth.Unmarshal()
   for _, test := range tests {
      var web WebAddress
      web.Set(test.url)
      content, err := auth.Content(web.Path)
      if err != nil {
         t.Fatal(err)
      }
      video, err := content.Video()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(encoding.Name(video))
      time.Sleep(time.Second)
   }
}
