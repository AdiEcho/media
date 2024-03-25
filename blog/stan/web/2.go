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
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.stan.com.au"
   req.URL.Path = "/login/v1/sessions/web/app"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader(url.Values{
   "jwToken":[]string{"eyJhbGciOiJIUzI1NiIsImtpZCI6InBpa2FjaHUiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE3MjE3NzIxNzQsImp0aSI6IjllZDRjNDE1ZTgyNzQ5NTg5MWJhZGMzOWNhZmNkZWEzIiwiaWF0IjoxNzExNDA0MTc0LCJyb2xlIjoidXNlciIsInVpZCI6ImUwNzUyOGZkM2I0NDRiMTQ4YTI0NmZmYjM5M2JlNjUyIiwic3RyZWFtcyI6ImhkIiwiY29uY3VycmVuY3kiOjMsInByb2ZpbGVJZCI6ImUwNzUyOGZkM2I0NDRiMTQ4YTI0NmZmYjM5M2JlNjUyIiwicHJvZmlsZU5hbWUiOiJzdGV2ZW4iLCJwaXIiOnRydWUsImFwcCI6IlN0YW4tV2ViIiwiZmVhdCI6MzM1NjM2MTk4NH0.l4Y0xxitjsdefXzxWKnA_3VvOKw83iQbAZjr0pvTsQI"},
}.Encode())
