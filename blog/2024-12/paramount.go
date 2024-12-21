package main

import (
   "net/http"
   "net/url"
   "os"
)

const content_id = "esJvFlqdrcS_kFHnpxSuYp449E7tTexD"

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.URL = &url.URL{}
   req.URL.Host = "www.paramountplus.com"
   req.URL.Path = "/apps-api/v2.0/androidphone/video/cid/"+content_id+".json"
   value := url.Values{}
   value["at"] = []string{"ABAAAAAAAAAAAAAAAAAAAAAA+c9AiUS4F1JOX0w0O1uzQw/+qdAuVgybB1FK7aqonjY="}
   req.URL.RawQuery = value.Encode()
   req.URL.Scheme = "https"
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   file, err := os.Create(content_id + ".json")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   file.ReadFrom(resp.Body)
}
