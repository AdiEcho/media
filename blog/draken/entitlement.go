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
   req.URL.Host = "client-api.magine.com"
   req.URL.Path = "/api/entitlement/v2/asset/17139357-ed0b-4a16-8be6-69e418c4ba40"
   req.URL.Scheme = "https"
   req.Header["Magine-Accesstoken"] = []string{"22cc71a2-8b77-4819-95b0-8c90f4cf5663"}
   req.Header["Authorization"] = []string{"Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwiaWF0IjoxNzE0ODY4NzExLCJ1c2VySWQiOiIxNTZLRUZKRExDN0hNSUtBVU5TUUJHOExTVVNSIiwidXNlckNvdW50cnkiOiJTRSIsIm5vZ2VvIjpmYWxzZSwiZGVidWciOmZhbHNlfQ.5Km1FYvwuj5J3aMepNxwzbjY3w7aXfWcWtRD9q3stwBC3nZgKSMOIJwhotQQF1r4SJqBKRC5sGoz7p77LDoKRA"}
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
