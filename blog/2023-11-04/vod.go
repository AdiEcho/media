package main

import (
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
)

/*
link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000283422?formats=m3u
link.theplatform.com/s/NnzsPC/media/guid/2410887629/9000283422?formats=mpeg-dash
*/
func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "lemonade.nbc.com"
   //pass
   req.URL.Path = "/v1/vod/2410887629/9000283422"
   //lock
   //req.URL.Path = "/v1/vod/2410887629/9000283426"
   req.URL.Scheme = "https"
   val := make(url.Values)
   val["browser"] = []string{"other"}
   // val["browser"] = []string{"safari"}
   //mpeg_cenc_2sec
   val["platform"] = []string{"web"}
   //mpeg_cenc
   //val["platform"] = []string{"android"}
   val["programmingType"] = []string{"Full Episode"}
   //val["programmingType"] = []string{"Clips"}
   req.URL.RawQuery = val.Encode()
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}
