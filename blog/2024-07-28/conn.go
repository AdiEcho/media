package main

import (
   "fmt"
   "net"
   "net/http"
   "os"
)

type connection struct {
   net.Conn
   raw []byte
}

func (c *connection) Read(text []byte) (int, error) {
   n, err := c.Conn.Read(text)
   c.raw = append(c.raw, text[:n]...)
   return n, err
}

func main() {
   req, err := http.NewRequest("", "http://example.com", nil)
   if err != nil {
      panic(err)
   }
   var (
      connect connection
      transport http.Transport
   )
   transport.Dial = func(network, addr string) (net.Conn, error) {
      connect.Conn, err = net.Dial(network, addr)
      if err != nil {
         return nil, err
      }
      return &connect, nil
   }
   resp, err := transport.RoundTrip(req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
   fmt.Printf("%q\n", connect.raw)
}
