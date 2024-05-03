package sbs

import (
   "net/http"
   "net/url"
   "os"
)

func video_stream() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.sbs.com.au"
   req.URL.Path = "/api/v3/video_stream"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["id"] = []string{"2229616195516"}
   val["context"] = []string{"odwebsite"}
   req.Header["Authorization"] = []string{"Bearer odwebsite9d30c7144aac68fbdf0caa716037371b17b9e149"}
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
