package http

import (
   "fmt"
   "io"
   "os"
   "testing"
)

func TestWrite(t *testing.T) {
   var resp response
   err := resp.New()
   if err != nil {
      t.Fatal(err)
   }
   body := resp.get_body()
   defer body.Close()
   file, err := os.Create("http.json")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   err = resp.set_body(io.TeeReader(body, file))
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", resp.Slideshow)
}

func TestRead(t *testing.T) {
   file, err := os.Open("http.json")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   var resp response
   err = resp.set_body(file)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", resp.Slideshow)
}
