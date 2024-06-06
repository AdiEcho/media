package rtbf

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
)

func (o one) three(login *accounts_login) (*http.Response, error) {
   var body = strings.NewReader(url.Values{
      "APIKey":[]string{api_key},
      // from /accounts.login
      "login_token":[]string{login.SessionInfo.CookieValue},
   }.Encode())
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"*/*"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "login.auvio.rtbf.be"
   req.URL.Path = "/accounts.getJWT"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.Header["Cookie"] = []string{
      // from /accounts.webSdkBootstrap
      "gmid=gmid.ver4.AtLtH4HMHg.LMhRVRJCFKP7uqs-cOeQLiHO5p4Gnf0AKg759MRJG72Xj9AzXsw20ySPPDaOmdSQ.EUz7cp0LCa8ATNMrSxDy9DuG5UvI5e_ZRJxrvDjrtEZJu-MqTqAcWHIz5ImHzxjzpS5i_tzQ8OOrRWUG07wvDg.sc3",
   }
   return http.DefaultClient.Do(&req)
}
