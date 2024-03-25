package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.stan.com.au"
   req.URL.Path = "/concurrency/v1/streams"
   req.URL.Scheme = "https"
   req.Header["x-forwarded-for"] = []string{"1.128.0.0"}
   val := make(url.Values)
   val["jwToken"] = []string{"eyJhbGciOiJIUzI1NiIsImtpZCI6InBpa2FjaHUiLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE3MjE3NzIxODUsImp0aSI6IjJiZThmYTBkYTg2NTQ0Njk4ZWUwMjg3YzdiZDc3YTIyIiwiaWF0IjoxNzExNDA0MTg1LCJyb2xlIjoidXNlciIsInVpZCI6ImUwNzUyOGZkM2I0NDRiMTQ4YTI0NmZmYjM5M2JlNjUyIiwic3RyZWFtcyI6ImhkIiwiY29uY3VycmVuY3kiOjMsInByb2ZpbGVJZCI6ImUwNzUyOGZkM2I0NDRiMTQ4YTI0NmZmYjM5M2JlNjUyIiwicHJvZmlsZU5hbWUiOiJzdGV2ZW4iLCJwaXIiOnRydWUsInR6IjoiQW1lcmljYS9DaGljYWdvIiwiYXBwIjoiU3Rhbi1XZWIiLCJ2ZXIiOiJiZWFkMDk2IiwiZmVhdCI6MzM1NjM2MTk4NH0.4A1MOC17P7bIA_hQhCowqhj1QSU-FJ5xJcyAktoASu4"}
   val["format"] = []string{"dash"}
   val["programId"] = []string{"1540676"}
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
