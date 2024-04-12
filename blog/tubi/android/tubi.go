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
   req.URL.Host = "uapi.adrise.tv"
   req.URL.Path = "/cms/content"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["content_id"] = []string{"589292"}
   val["deviceId"] = []string{"ab55452c-66e0-4021-9619-5bdc25f26ae8"}
   val["platform"] = []string{"android"}
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
