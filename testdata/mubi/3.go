package main

import (
   "net/http"
   "net/url"
   "os"
)

const bearer = "79df72c589f824c61245ba11ff0a4c63fb13a1"

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.mubi.com"
   req.URL.Path = "/v3/films/187/viewing/secure_url"
   req.URL.Scheme = "https"
   req.Header["Client"] = []string{"web"}
   req.Header["Client-Country"] = []string{"GR"}
   req.Header["Authorization"] = []string{"Bearer " + bearer}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
