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
   req.Header["Authorization"] = []string{"Bearer eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi01Mjg2NDYyMi1hYjM0LTQzZWUtODFjMi0yMDUwZWYzZjRmNzQiLCJpc3MiOiJmcGEtaXNzdWVyIiwic3ViIjoiVVNFUklEOmJvbHQ6YWQwNmY5N2ItMTk0NC00NDg5LTlmM2UtNTRjMmEwNTcyMzdlIiwiaWF0IjoxNzE3OTYzMjA3LCJleHAiOjIwMzMzMjMyMDcsInR5cGUiOiJBQ0NFU1NfVE9LRU4iLCJzdWJkaXZpc2lvbiI6ImJlYW1fYW1lciIsInNjb3BlIjoiZGVmYXVsdCIsInZlcnNpb24iOiJ2MiIsImFub255bW91cyI6dHJ1ZSwiZGV2aWNlSWQiOiI1MjlmYTA3MGJjNzBhNjcwIn0.ZsdFpYkEToGHCg4bAJLv_WT5Sl3xz42NrTWN124MLbesDCLNC0XXasGq3Sq4H_rI8JAf5CzL76Q8khWjXuzBwEN_ihqrObK8i30xv32UNjjZwt5CqhGnutbwSbV4jvXqRx7oH3wlSWTXRtuquZw5tnZj0EovFGAV7yFZ78NM-BI1VJtdSEgxnEktqG83ZVU2vO7L78vwtRkemKSaHwKDBmrECzwb8GTN6BUrAkxOKAadDX3pO6ayOlL4AjmsQEF5QXiTQNl-Kx7EYJwrh8T7Y4Dsou-u6pLINmUfxcK0pKR2QdK4dIkCPFIB5NxBBxYMlv_1hob4r9BGkSR0wqop9g"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["X-Disco-Arkose-Token"] = []string{"36617d76e77127a81.2715576401|r=us-east-1|meta=3|meta_width=300|metabgclr=transparent|metaiconclr=%23555555|guitextcolor=%23000000|lang=en|pk=B0217B00-2CA4-41CC-925D-1EEB57BFFC2F|at=40|sup=1|rid=25|ag=101|cdn_url=https%3A%2F%2Fwbd-api.arkoselabs.com%2Fcdn%2Ffc|lurl=https%3A%2F%2Faudio-us-east-1.arkoselabs.com|surl=https%3A%2F%2Fwbd-api.arkoselabs.com|smurl=https%3A%2F%2Fwbd-api.arkoselabs.com%2Fcdn%2Ffc%2Fassets%2Fstyle-manager"}
   req.Header["X-Disco-Client-Id"] = []string{"android1_prd:1717963242:d1add02cacecf708d283283514ef5385ee195e29a8c18f8c1579a99720f2404d"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "default.any-amer.prd.api.discomax.com"
   req.URL.Path = "/login"
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
 "credentials": {
  "username": "EMAIL",
  "password": "PASSWORD"
 }
}`)
