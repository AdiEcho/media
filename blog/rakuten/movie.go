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
   req.URL.Host = "gizmo.rakuten.tv"
   req.URL.Path = "/v3/movies/jerry-maguire"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["classification_id"] = []string{"23"}
   val["market_code"] = []string{"fr"}
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
