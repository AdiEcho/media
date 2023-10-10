package main

import (
   "fmt"
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
   req.URL.Host = "www.youtube.com"
   req.URL.Path = "/youtubei/v1/player"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
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
   if strings.Contains(string(res_body), "This video requires payment to watch") {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}

var req_body = strings.NewReader(`
{
   "videoId": "oCjW6gdEDa4",
   "context": {
      "client": {
         "clientName": "IOS",
         "clientVersion": "17.33.2"
      }
   }
}
`)
