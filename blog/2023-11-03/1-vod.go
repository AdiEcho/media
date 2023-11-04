package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "lemonade.nbc.com"
   req.URL.Path = "/v1/vod/2410887629/9000283422"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["programmingType"] = []string{"Full Episode"}
   val["platform"] = []string{"web"}
   //val["platform"] = []string{"android"}
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
