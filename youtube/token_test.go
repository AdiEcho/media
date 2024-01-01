package youtube

import (
   "os"
   "testing"
   "time"
)

var check_ids = []string{
   "Cr381pDsSsA", // racy check
   "HsUATh_Nc2U", // racy check
   "SZJvDhaSDnc", // racy check
   "Tq92D6wQ1mg", // racy check
   "dqRZDebPIGs", // racy check
   "nGC3D_FkCmg", // content check
}

func Test_Android_Check(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   raw, err := os.ReadFile(home + "/youtube.json")
   if err != nil {
      t.Fatal(err)
   }
   for _, check_id := range check_ids {
      var play Player
      req := Request{Video_ID: check_id}
      req.Android_Check()
      var tok Token
      tok.Unmarshal(raw)
      err := play.Post(req, &tok)
      if err != nil {
         t.Fatal(err)
      }
      if play.Playability.Status != "OK" {
         t.Fatal(play)
      }
      time.Sleep(time.Second)
   }
}
