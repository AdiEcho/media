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
   req.Header["Accept"] = []string{"application/json"}
   req.Header["Accept-Encoding"] = []string{"gzip, deflate, br"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Content-Length"] = []string{"912"}
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Origin"] = []string{"https://auvio.rtbf.be"}
   req.Header["Referer"] = []string{"https://auvio.rtbf.be/"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"cross-site"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"}
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "exposure.api.redbee.live"
   req.URL.Path = "/v2/customer/RTBF/businessunit/Auvio/auth/gigyaLogin"
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
 "jwt": "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IlJFUTBNVVE1TjBOQ1JUSkVNemszTTBVMVJrTkRRMFUwUTBNMVJFRkJSamhETWpkRU5VRkJRZyJ9.eyJpc3MiOiJodHRwczovL2ZpZG0uZ2lneWEuY29tL2p3dC80X01sX2ZKNDdHbkJBVzZGclB6TXhoMHcvIiwiYXBpS2V5IjoiNF9NbF9mSjQ3R25CQVc2RnJQek14aDB3IiwiaWF0IjoxNzE3NjM0Mzk0LCJleHAiOjE3MTc2MzQ2OTQsInJpc2tTY29yZSI6MC4wLCJzdWIiOiI4OTgyNjk5MDUwZWQ0MWY5OGZlZDA4NzdiZGE2ZjYxNiIsImVtYWlsIjoiMjAyNC02LTVAbWFpbHNhYy5jb20iLCJmaXJzdE5hbWUiOiJzdGV2ZW4iLCJsYXN0TmFtZSI6InBlbm55In0.MH7dkz4TuVu3RcqblnHBYBNMd2908uKBSUmFP6gGlMw3Of8HiYrHIVlz5zxhTX73y60RjjnvZJg_R6v5ovDwY2IE4JZVwCkgdW2_qeuarq_Sj-NJpGufzLY9iop2Q8cE3qXo9Sl20O3yDU852WXurTHr4WX-weO6bHAz-lDB1QH4nbzTM0qM4_IhV2iToa_NNbDhutYLvxJkPBtgfKm10xHMUmmM6dsmcr_tF2Ge4r_W63WiogioNmZG5Q4pZig9-B9dXFeJ0onI5cTTLBk1lzomvxNTG9ugxJOTRdUb7JMgx9SMoYsZTCuExcq4xWrQfVigqEs5l0Vh47LDQ5C-OA",
 "device": {
  "deviceId": "7f5cdd55-1cfe-4841-9e8e-ecd8b823cfad",
  "name": "Browser",
  "type": "WEB"
 }
}`)
