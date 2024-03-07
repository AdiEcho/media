package web

import (
   "os"
   "testing"
)

const file_id = "392482fe9bed7372d1657d7e22f32b792902f3bd"

func TestStorage(t *testing.T) {
   res, err := storage(file_id)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
