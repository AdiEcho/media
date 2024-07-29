package main

import (
   "crypto/tls"
   "net/http"
   "os"
)

func main() {
   req, err := http.NewRequest("", "https://example.com", nil)
   if err != nil {
      panic(err)
   }
   req.Close = true
   conn, err := tls.Dial("tcp", "example.com:https", nil)
   if err != nil {
      panic(err)
   }
   err = req.Write(conn)
   if err != nil {
      panic(err)
   }
   os.Stdout.ReadFrom(conn)
}
