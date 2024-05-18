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
   req.URL.Host = "auth.vhx.com"
   req.URL.Path = "/v1/oauth/token"
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
   "password":[]string{"PASSWORD"},
   "username":[]string{"USERNAME"},
   "client_id":[]string{"9a87f110f79cd25250f6c7f3a6ec8b9851063ca156dae493bf362a7faf146c78"},
   "grant_type":[]string{"password"},
}.Encode())
