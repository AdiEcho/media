package main

import (
   "io"
   "net/http"
   "net/url"
   "os"
   "strings"
   "fmt"
)

const auth_token = "faf0cd0e109db2eb320a1846f6547562"

func main() {
   body := fmt.Sprintf(`
   {"auth_token": %q}
   `, auth_token)
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "api.mubi.com"
   req.URL.Path = "/v3/authenticate"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(strings.NewReader(body))
   req.Header["Accept"] = []string{"application/json"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Client"] = []string{"android"}
   req.Header["Client-Accept-Audio-Codecs"] = []string{"AAC"}
   req.Header["Client-App"] = []string{"mubi"}
   req.Header["Client-Country"] = []string{"US"}
   req.Header["Client-Device-Brand"] = []string{"unknown"}
   req.Header["Client-Device-Identifier"] = []string{"437cacfa-7421-410c-a4cd-fbe5d5460345"}
   req.Header["Client-Device-Model"] = []string{"Android SDK built for x86"}
   req.Header["Client-Device-Os"] = []string{"6.0"}
   req.Header["Client-Version"] = []string{"41.2"}
   req.Header["Content-Length"] = []string{"49"}
   req.Header["Content-Type"] = []string{"application/json; charset=UTF-8"}
   req.Header["User-Agent"] = []string{"okhttp/4.10.0"}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
