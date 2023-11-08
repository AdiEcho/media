package main

import (
   "encoding/json"
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var u struct { Username, Password string }
   {
      s, err := os.UserHomeDir()
      if err != nil {
         panic(err)
      }
      b, err := os.ReadFile(s + "/hulu.json")
      if err != nil {
         panic(err)
      }
      json.Unmarshal(b, &u)
   }
   var req_body = strings.NewReader(url.Values{
      "friendly_name":[]string{"!"},
      "password":[]string{u.Password},
      "serial_number":[]string{"!"},
      "user_email":[]string{u.Username},
   }.Encode())
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "auth.hulu.com"
   // github.com/matthuisman/slyguy.addons/blob/master/slyguy.hulu/resources/lib/api.py
   req.URL.Path = "/v2/livingroom/password/authenticate"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(req_body)
   req.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(res_body)
}
