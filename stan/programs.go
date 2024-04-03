package stan

import (
   "net/http"
   "net/url"
)

func programs() (*http.Response, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.stan.com.au"
   req.URL.Scheme = "https"
   req.URL.Path = "/programs/v1/legacy/programs/1768588.json"
   return http.DefaultClient.Do(&req)
}
