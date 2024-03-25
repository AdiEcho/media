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
   req.URL.Host = "api.stan.com.au"
   req.URL.Path = "/login/v1/sessions/web/app"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader(url.Values{
   "jwToken":[]string{"eyJhbGciOiJIUzI1NiIsImtpZCI6InBpa2FjaHUiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE3MjE3NzIxODMsImp0aSI6IjE4N2RhNDdhMTM4ZjQ3OTY5NjMzNTUwYTcyOWIwODY0IiwiaWF0IjoxNzExNDA0MTgzLCJyb2xlIjoidXNlciIsInVpZCI6ImUwNzUyOGZkM2I0NDRiMTQ4YTI0NmZmYjM5M2JlNjUyIiwic3RyZWFtcyI6ImhkIiwiY29uY3VycmVuY3kiOjMsInByb2ZpbGVJZCI6ImUwNzUyOGZkM2I0NDRiMTQ4YTI0NmZmYjM5M2JlNjUyIiwicHJvZmlsZU5hbWUiOiJzdGV2ZW4iLCJwaXIiOnRydWUsInR6IjoiQW1lcmljYS9DaGljYWdvIiwiYXBwIjoiU3Rhbi1XZWIiLCJ2ZXIiOiJiZWFkMDk2IiwiZmVhdCI6MzM1NjM2MTk4NH0.ZUARuf7IvBBDRsbNJxjhGN4AnVTnarF0t-ZGTv_TFjI"},
}.Encode())
