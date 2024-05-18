package main

import (
   "io"
   "net/http"
   "net/url"
   "strings"
   "fmt"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.criterionchannel.com"
   req.URL.Path = "/login"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.Header["Cookie"] = []string{
      "_session=TEpHYVpYZ0M5dFVlQ0MzSVVtdVRJb1o3RTBFa2tOVWd3TUJaajA0THFUQ2svS09EUS82MDE1dnhacFVGQmpLYzUzZzFaWStoSnVaajJacnIwTGFucjdwZHU5bGpjU2EzK1REaVpNbFB6VGNCeEwyd3orenliejB2VVpyZXRSK3dSWE40NnNGS3lnR1p6eFBaQ2FwWHBTSm1ZR0gwK08reHZHNGZkaVpLUnNJNlQ2UUl2eHl3K0FsWkVBREdPZ1FSLS11endNQWlsY1Vqeld2VVFjSm55L2RRPT0%3D--2a64dbf1e5a4b48d3de0c3483692434fb59c9814",
   }
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   fmt.Printf("%+v\n", res)
}

var body = strings.NewReader(url.Values{
   "email":[]string{"EMAIL"}, "password":[]string{"PASSWORD"},
   "authenticity_token":[]string{"evo3lO59MN7AGAbYZMzCPYHn7Cvmr8Gw/KNge8wLgZ5KYD3DC6jKGdbgHuxRYvLWuHy+zAXOR6HzcsHaTvgnsQ=="},
}.Encode())
