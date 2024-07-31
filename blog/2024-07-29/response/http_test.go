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
   os.WriteFile("header.txt", []byte(resp.fly_request_id), 0666)
   os.WriteFile("body.json", resp.raw_body, 0666)
   err = resp.unmarshal_body()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", resp.body)
}

func TestRead(t *testing.T) {
   var resp response
   raw, err := os.ReadFile("header.txt")
   if err != nil {
      t.Fatal(err)
   }
   resp.fly_request_id = string(raw)
   resp.raw_body, err = os.ReadFile("body.json")
   if err != nil {
      t.Fatal(err)
   }
   err = resp.unmarshal_body()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", resp.body)
}
