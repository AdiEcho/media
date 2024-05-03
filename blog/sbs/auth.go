package sbs

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func auth() {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "www.sbs.com.au"
   req.URL.Path = "/api/v3/janrain/auth_native_traditional"
   req.URL.Scheme = "https"
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   body := url.Values{
      "pass": {"PASS"},
      "user": {"USER"},
   }.Encode()
   req.Body = io.NopCloser(strings.NewReader(body))
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
