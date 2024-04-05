package plex

import (
   "os"
   "testing"
)

func TestAnonymous(t *testing.T) {
   res, err := anonymous()
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
