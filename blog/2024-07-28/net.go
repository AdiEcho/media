package main

import (
   "net"
   "net/http"
   "os"
)

func main() {
   conn, err := net.Dial("tcp", "example.com:80")
   if err != nil {
      panic(err)
   }
   req, err := http.NewRequest("GET", "http://example.com", nil)
   if err != nil {
      panic(err)
   }
   req.Close = true
   req.Write(conn)
   os.Stdout.ReadFrom(conn)
}
