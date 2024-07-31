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
   os.WriteFile("http.json", resp.raw, 0666)
   err = resp.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", resp.body)
}

func TestRead(t *testing.T) {
   var (
      resp response
      err error
   )
   resp.raw, err = os.ReadFile("http.json")
   if err != nil {
      t.Fatal(err)
   }
   err = resp.unmarshal()
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", resp.body)
}
