package http

import (
   "fmt"
   "os"
   "testing"
)

func TestWrite(t *testing.T) {
   var body response_body
   err := body.New()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("http.json", body.get(), 0666)
   err = body.set(body.get())
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", body.Slideshow)
}

func TestRead(t *testing.T) {
   raw, err := os.ReadFile("http.json")
   if err != nil {
      t.Fatal(err)
   }
   var body response_body
   err = body.set(raw)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", body.Slideshow)
}
