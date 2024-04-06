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
   req.URL.Host = "discover.provider.plex.tv"
   req.URL.Path = "/library/metadata/matches"
   req.URL.Scheme = "http"
   req.Header["Accept"] = []string{"application/json"}
   val := make(url.Values)
   //val["url"] = []string{"/movie/cruel-intentions"}
   val["url"] = []string{"/show/broadchurch/season/3/episode/5"}
   val["X-Plex-Token"] = []string{"aREUTWtbGNN8p_ChaGpv"}
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
