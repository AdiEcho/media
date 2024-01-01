package amc

import (
   "154.pages.dev/stream"
   "fmt"
   "os"
   "testing"
   "time"
)

func Test_Content(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   raw, err := os.ReadFile(home + "/amc/auth.json")
   if err != nil {
      t.Fatal(err)
   }
   auth, err := Raw_Auth.Unmarshal(raw)
   if err != nil {
      t.Fatal(err)
   }
   for _, test := range tests {
      con, err := auth.Content(test.u)
      if err != nil {
         t.Fatal(err)
      }
      vid, err := con.Video()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(stream.Name(vid))
      time.Sleep(time.Second)
   }
}
