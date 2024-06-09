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
   req.Header["X-Disco-Arkose-Token"] = []string{"62717d76c76330fb6.3222074501|r=us-east-1|meta=3|meta_width=300|metabgclr=transparent|metaiconclr=%23555555|guitextcolor=%23000000|lang=en|pk=B0217B00-2CA4-41CC-925D-1EEB57BFFC2F|at=40|sup=1|rid=84|ag=101|cdn_url=https%3A%2F%2Fwbd-api.arkoselabs.com%2Fcdn%2Ffc|lurl=https%3A%2F%2Faudio-us-east-1.arkoselabs.com|surl=https%3A%2F%2Fwbd-api.arkoselabs.com|smurl=https%3A%2F%2Fwbd-api.arkoselabs.com%2Fcdn%2Ffc%2Fassets%2Fstyle-manager"}
   req.Header["X-Disco-Client-Id"] = []string{"web1_prd:1717961040:a73067989c7beb6d18fab97d596fb6448aa990f4c2e36cdb29a89f1da36f6ec4"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "default.any-amer.prd.api.max.com"
   req.URL.Path = "/login"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Cookie"] = []string{
      "st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi02ZTA1OGEwMi1iMmIxLTQ3MzAtOWY4ZC1jZWRjZjg4Yjg2NTAiLCJpc3MiOiJmcGEtaXNzdWVyIiwic3ViIjoiVVNFUklEOmJvbHQ6NzI4Y2Q4Y2UtZWNmYi00Mjc5LWExNTItNmE1NzI2MjU4OGM1IiwiaWF0IjoxNzE3OTYxMDI3LCJleHAiOjIwMzMzMjEwMjcsInR5cGUiOiJBQ0NFU1NfVE9LRU4iLCJzdWJkaXZpc2lvbiI6ImJlYW1fYW1lciIsInNjb3BlIjoiZGVmYXVsdCIsInZlcnNpb24iOiJ2MiIsImFub255bW91cyI6dHJ1ZSwiZGV2aWNlSWQiOiIzNDVkY2MyNi03ZTljLTRjYmUtOTU4MS1hOTVmZDk0MjdhN2IifQ.X8CIJqou5HMtmcHScESraeHBzqqAc8rY6chX7cJDfJEuovuVbLh3G4hWa1NetR3XZaK8sUyrg9RsOGHeXWjXT03Z0g4TParqZ5k90FNc4OZF_Q_7aVgVqhlTeBoIBWQ2O6irlgrap_p9UYXcRy9kC7xMKmeqa6tZ0j2WKXYc_G7mj5iiNweoXvZ7G4bRkVjsmVBnyxpxTbnYEMN-E_lfbL4h0-bV_qhrDZIzvMnmoUqZVBHgxjidwmBttyySv6T38XZCjPDJgwVbSfQ08CvYpQMSUOQgizHb44pK8Q6i6r0Kh8vKc1yvscr-7sd6Dc2rkNithGiQk1WbepWHHpem8w",
   }
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader(`
{
 "credentials": {
  "username": "EMAIL",
  "password": "PASSWORD"
 }
}
`)
