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
   req.Header["Accept"] = []string{"application/json, text/plain, */*"}
   req.Header["Accept-Language"] = []string{"nl"}
   req.Header["Connection"] = []string{"keep-alive"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Host"] = []string{"api.audienceplayer.com"}
   req.Header["Origin"] = []string{"https://www.cinemember.nl"}
   req.Header["Referer"] = []string{"https://www.cinemember.nl/"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"cross-site"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.audienceplayer.com"
   req.URL.Path = "/graphql/2/user"
   req.URL.RawPath = ""
   val := make(url.Values)
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader(`{
 "operationName": "UserAuthenticate",
 "variables": {
  "email": "EMAIL",
  "password": "PASSWORD"
 },
 "query": "mutation UserAuthenticate($email: String, $password: String) {\n  UserAuthenticate(email: $email, password: $password) {\n    access_token\n  }\n}"
}
`)
