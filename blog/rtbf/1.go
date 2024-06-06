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
   req.URL.Host = "login.auvio.rtbf.be"
   req.URL.Path = "/accounts.login"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   req.Header["Cookie"] = []string{
      "gmid=gmid.ver4.AtLtH4HMHg.LMhRVRJCFKP7uqs-cOeQLiHO5p4Gnf0AKg759MRJG72Xj9AzXsw20ySPPDaOmdSQ.EUz7cp0LCa8ATNMrSxDy9DuG5UvI5e_ZRJxrvDjrtEZJu-MqTqAcWHIz5ImHzxjzpS5i_tzQ8OOrRWUG07wvDg.sc3",
   }
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader(url.Values{
   "APIKey":[]string{"4_Ml_fJ47GnBAW6FrPzMxh0w"},
   "loginID":[]string{"EMAIL"},
   "password":[]string{"PASSWORD"},
}.Encode())
