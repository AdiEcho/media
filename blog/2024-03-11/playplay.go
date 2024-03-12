package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "guc3-spclient.spotify.com"
   req.URL.Path = "/playplay/v1/key/f682d2a95d0e14eeef4f40b60fddde56bc6721c7"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Authorization"] = []string{"Bearer BQCHAsDof0SnWGNOSmBjfxXr0cC_eqdLY8N_fL_54XlqESZOQ2qZ0q65s4TWOIMXDXF1siJgUnfZKAei44OSIuIXZovqA2go5byZ5MEFYshZYlUT3Jh2EKw2wIRyaXhjlaAxtqizrr4uF0izrNYPpCzk61Um1idTuJXwTVSMxkXrdWEN9RcNMh0nSZzCmhCkDtcvokHGSd3MSK7QiwFad_d80FaoG2b4PivFzfllNic6a6m4F4t22ztD9Ho70I8Tb30M5ewyS5aztG5BF29rZiNqNdvAoS5evDSLOuoz32mKMM_CNV7uGSyaySW-W0C9m9B1Knw19JKhDyH2NX6_"}
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader("\b\x02\x12\x10\x01K\xe0K\xce^\xe6\xb3nl\xec0\xd8\xeb\x9a2 \x01(\x010\xfa\x92\x85\xaf\x06")
