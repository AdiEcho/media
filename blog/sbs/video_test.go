package sbs

import (
   "os"
   "testing"
)

func TestVideo(t *testing.T) {
   user, pass := os.Getenv("sbs_username"), os.Getenv("sbs_password")
   if user == "" {
      t.Fatal("Getenv")
   }
   var auth auth_native
   err := auth.New(user, pass)
   if err != nil {
      t.Fatal(err)
   }
   res, err := auth.video_stream()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
