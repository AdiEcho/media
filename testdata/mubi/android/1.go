package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.mubi.com"
   req.URL.Path = "/v3/link_code"
   req.URL.Scheme = "https"
   req.Header["Client"] = []string{"android"}
   req.Header["Client-Country"] = []string{"US"}
   req.Header["Client-Device-Identifier"] = []string{"!"}
   req.Header["Client-Version"] = []string{"!"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
