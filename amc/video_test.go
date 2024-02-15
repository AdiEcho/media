package amc

import (
   "154.pages.dev/rosso"
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
   raw, err := os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth, err := RawAuth.Unmarshal(raw)
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      content, err := auth.Content(test.u)
      if err != nil {
         t.Fatal(err)
      }
      video, err := content.Video()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(rosso.Name(video))
      time.Sleep(time.Second)
   }
}
