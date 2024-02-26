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
   req.URL.Host = "atom.peacocktv.com"
   req.URL.Path = "/adapter-calypso/v3/query/node/content_id/GMO_00000000224510_02_HDSDR"
   req.URL.Scheme = "https"
   req.Header["X-Skyott-Proposition"] = []string{"NBCUOTT"}
   req.Header["X-Skyott-Territory"] = []string{"US"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
