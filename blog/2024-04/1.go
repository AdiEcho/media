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
   req.URL.Host = "sas.peacocktv.com"
   req.URL.Path = "/companion-service/journeys/sign-in"
   req.URL.Scheme = "https"
   req.Header["X-Skyott-Device"] = []string{"TV"}
   req.Header["X-Skyott-Platform"] = []string{"ANDROIDTV"}
   req.Header["X-Skyott-Proposition"] = []string{"NBCUOTT"}
   req.Header["X-Skyott-Provider"] = []string{"NBCU"}
   req.Header["X-Skyott-Territory"] = []string{"US"}
   req.Header["Content-Type"] = []string{"application/vnd.companionservice.v1+json"}
   body := strings.NewReader(`
   {"deviceId":""}
   `)
   req.Body = io.NopCloser(body)
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

