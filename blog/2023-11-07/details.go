package main

import (
   "io"
   "net/http"
   "net/http/httputil"
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
   req.URL.Host = "guide.hulu.com"
   req.URL.Path = "/guide/details"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
   val := make(url.Values)
   val["user_token"] = []string{"I4YwKXcuqJuV_9OPFdUo6zHpW7w-GXb0ngpaYieW8dv..."}
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}

var req_body = strings.NewReader(`
{
  "eabs": [
    "EAB::023c49bf-6a99-4c67-851c-4c9e7609cc1d::196861183::262714326"
  ]
}
`)
