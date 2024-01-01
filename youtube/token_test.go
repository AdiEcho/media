package youtube

import (
   "os"
   "testing"
   "time"
)

func Test_Android_Check(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   tok, err := Read_Token(home + "/youtube.json")
   if err != nil {
      t.Fatal(err)
   }
   var req Request
   req.Android_Check()
   for _, check := range check_IDs {
      req.Video_ID = check
      play, err := req.Player(tok)
      if err != nil {
         t.Fatal(err)
      }
      if play.Playability.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}
