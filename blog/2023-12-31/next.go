package main

import (
   "fmt"
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
   req.URL.Host = "www.youtube.com"
   req.URL.Scheme = "https"
   req.URL.RawQuery = "prettyPrint=false" // default is pretty
   req.URL.Path = "/youtubei/v1/next"
   req.Body = func() io.ReadCloser {
      s := strings.NewReader(`
      {
         "videoId": "2ZcDwdXEVyI",
         "context": {
            "client": {
               "clientName": "WEB",
               "clientVersion": "2.20231219.04.00"
            }
         }
      }
      `)
      return io.NopCloser(s)
   }()
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      panic(res.Status)
   }
   body, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   lower := strings.ToLower(string(body))
   if strings.Contains(lower, "in the heat of the night") {
      fmt.Println("pass", len(body))
      os.WriteFile("WEB.json", body, 0666)
   } else {
      fmt.Println("fail", len(body))
   }
}
