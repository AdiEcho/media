package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.joyn.de"
   req.URL.Path = "/graphql"
   val := make(url.Values)
   req.URL.Scheme = "https"
   req.Header["Joyn-Platform"] = []string{"web"}
   req.Header["X-Api-Key"] = []string{"4f0fd9f18abbe3cf0e87fdb556bc39c8"}
   val["variables"] = []string{"{\"path\":\"/filme/barry-seal-only-in-america\"}"}
   val["extensions"] = []string{"{\"persistedQuery\":{\"version\":1,\"sha256Hash\":\"5cd6d962be007c782b5049ec7077dd446b334f14461423a72baf34df294d11b2\"}}"}
   // optional:
   val["operationName"] = []string{"PageMovieDetailStatic"}
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
