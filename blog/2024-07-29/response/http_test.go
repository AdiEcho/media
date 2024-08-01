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
   os.WriteFile("date.txt", []byte(resp.raw.date), 0666)
   os.WriteFile("body.txt", resp.raw.body, 0666)
   err = resp.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(resp.date)
   fmt.Printf("%+v\n", resp.body)
}

func TestRead(t *testing.T) {
   var resp response
   // date
   raw, err := os.ReadFile("date.txt")
   if err != nil {
      t.Fatal(err)
   }
   resp.raw.date = string(raw)
   // body
   resp.raw.body, err = os.ReadFile("body.txt")
   if err != nil {
      t.Fatal(err)
   }
   err = resp.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Println(resp.date)
   fmt.Printf("%+v\n", resp.body)
}
