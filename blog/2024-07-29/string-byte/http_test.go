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
   os.WriteFile("date.txt", []byte(resp.date.raw), 0666)
   os.WriteFile("body.txt", resp.body.raw, 0666)
   err = resp.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(resp.date.value)
   fmt.Printf("%+v\n", resp.body.value)
}

func TestRead(t *testing.T) {
   var resp response
   raw, err := os.ReadFile("date.txt")
   if err != nil {
      t.Fatal(err)
   }
   resp.date.raw = string(raw)
   resp.body.raw, err = os.ReadFile("body.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = resp.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(resp.date.value)
   fmt.Printf("%+v\n", resp.body.value)
}
