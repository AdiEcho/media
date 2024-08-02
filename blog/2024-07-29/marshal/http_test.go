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
   text, err := resp.marshal()
   if err != nil {
      t.Fatal(err)
   }
   os.WriteFile("response.json", text, 0666)
   fmt.Printf("%+v\n", resp)
}

func TestRead(t *testing.T) {
   text, err := os.ReadFile("response.json")
   if err != nil {
      t.Fatal(err)
   }
   var resp response
   err = resp.unmarshal(text)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", resp)
}
