package sbs

import (
   "net/http"
   "net/url"
   "os"
   "strings"
)

func streams() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "pubads.g.doubleclick.net"
   req.URL.Path = "/ondemand/hls/content/2488267/vid/2229616195516A/streams"
   req.URL.Scheme = "https"
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
