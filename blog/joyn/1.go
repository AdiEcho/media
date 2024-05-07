package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "auth.joyn.de"
   req.URL.Path = "/auth/anonymous"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/json"}
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader(`
{
   "client_name": "web",
   "client_id": "!"
}
`)
