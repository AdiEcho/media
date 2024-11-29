package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "www.kanopy.com"
   req.URL.Path = "/kapi/memberships"
   value := url.Values{}
   value["userId"] = []string{"8177465"}
   req.URL.RawQuery = value.Encode()
   req.URL.Scheme = "https"
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVpZCI6IjgxNzc0NjUiLCJpZGVudGl0eV9pZCI6IjQxOTQ1NjkyOSIsInZpc2l0b3JfaWQiOiIxNzMyODIzODAzOTUzMDMxNzE5Iiwic2Vzc2lvbl9pZCI6IjE3MzI4MjM4MDM5NTMwODc0MDMiLCJjb25uZWN0aW9uX2lkIjoiMTczMjgyMzgwMzk1MzA4NzQwMyIsImt1aV91c2VyIjoxLCJyb2xlcyI6WyJjb21Vc2VyIl19LCJpYXQiOjE3MzI4MjM4MDMsImV4cCI6MjA0ODE4MzgwMywiaXNzIjoia2FwaSJ9.M6n5KPLzsLE1U8xuWc1tQ_gCAIUCFb4BtJXQDHd07m8"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/115.0"}
   req.Header["X-Version"] = []string{"web/prod/4.16.0/2024-11-07-14-23-23"}
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}
