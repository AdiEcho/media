package sbs

import (
   "os"
   "testing"
)

func TestAuth(t *testing.T) {
   user, pass := os.Getenv("sbs_username"), os.Getenv("sbs_password")
   if user == "" {
      t.Fatal("Getenv")
   }
   res, err := auth_native(user, pass)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
