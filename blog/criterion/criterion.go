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
   req.URL.Host = "embed.criterionchannel.com"
   req.URL.Path = "/videos/455774"
   req.URL.Scheme = "https"
   req.Header["Referer"] = []string{"https://www.criterionchannel.com/"}
   val := make(url.Values)
   val["auth-user-token"] = []string{"eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjo2NTc0NjE0NiwiZXhwIjoxNzE1OTkzMDAzfQ.OFddwkdAvKjriRtm07-9l-2Jn_j4f1_IsdeNKfr3ntE"}
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
