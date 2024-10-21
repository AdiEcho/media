package max

import (
   "net/http"
   "net/url"
)

// you must
// /authentication/linkDevice/initiate
// first or this will always fail
func (b bolt_token) login() (*http.Response, error) {
   var req http.Request
   req.Header = http.Header{}
   req.Method = "POST"
   req.URL = &url.URL{}
   req.URL.Host = "default.prd.api.discomax.com"
   req.URL.Path = "/authentication/linkDevice/login"
   req.URL.Scheme = "https"
   req.AddCookie(b.st)
   return http.DefaultClient.Do(&req)
}
