package sbs

import (
   "fmt"
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
   video, err := auth.video()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", video)
}
