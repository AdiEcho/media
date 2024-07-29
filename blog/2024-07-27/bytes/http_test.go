package http

import (
   "fmt"
   "os"
   "testing"
)

func TestWrite(t *testing.T) {
   var resp response
   err := resp.New()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("http.json", resp.marshal(), 0666)
   err = resp.unmarshal(resp.marshal())
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", resp.Slideshow)
}

func TestRead(t *testing.T) {
   raw, err := os.ReadFile("http.json")
   if err != nil {
      t.Fatal(err)
   }
   var resp response
   err = resp.unmarshal(raw)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", resp.Slideshow)
}
