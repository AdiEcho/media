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
   req.URL.Path = "/v3/films/187/viewing/secure_url"
   req.URL.Scheme = "https"
   req.Header["Client"] = []string{"web"}
   req.Header["Client-Country"] = []string{"GR"}
   req.Header["Authorization"] = []string{"Bearer 1abf8440609dc3b9c835a9c8b1445319fb13a1"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
