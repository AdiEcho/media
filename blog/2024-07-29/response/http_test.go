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
   os.WriteFile("header.txt", []byte(resp.header.fly_request_id), 0666)
   os.WriteFile("body.json", resp.body.raw, 0666)
   err = resp.set_body(resp.get_body())
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", resp.body.Slideshow)
}

func TestRead(t *testing.T) {
   var resp response
   raw, err := os.ReadFile("header.txt")
   if err != nil {
      t.Fatal(err)
   }
   resp.header.fly_request_id = string(raw)
   raw, err = os.ReadFile("body.json")
   if err != nil {
      t.Fatal(err)
   }
   err = resp.set_body(raw)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", resp.body.Slideshow)
}
