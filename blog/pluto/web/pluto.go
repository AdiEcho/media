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
   req.URL.Host = "boot.pluto.tv"
   req.URL.Path = "/v4/start"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["appName"] = []string{"web"}
   val["appVersion"] = []string{"9"}
   val["clientID"] = []string{"9"}
   val["clientModelNumber"] = []string{"9"}
   val["episodeSlugs"] = []string{"ex-machina-2015-1-1-ptv1"}
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
