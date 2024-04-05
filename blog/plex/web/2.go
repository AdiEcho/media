package plex

import (
   "net/http"
   "net/url"
   "os"
)

func metadata() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "vod.provider.plex.tv"
   req.URL.Path = "/library/metadata/movie:cruel-intentions"
   req.URL.Scheme = "https"
   req.Header["X-Plex-Token"] = []string{"fc1WPqnLdmq3J4Axt5pn"}
   req.Header["Accept"] = []string{"application/json"}
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
