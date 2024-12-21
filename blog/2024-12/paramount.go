package main

import (
   "net/http"
   "net/url"
   "os"
)

const (
   at = "ABAAAAAAAAAAAAAAAAAAAAAAU/Emq2DGAxNGYi71fZdi2mb6UTU2+elcD3bGEhs6bo4="
   content_id = "ssc3CuuS4mrQ7EyVXILH0FEQSi5yBAsA"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.URL = &url.URL{}
   req.URL.Host = "www.paramountplus.com"
   req.URL.Path = "/apps-api/v2.0/androidphone/video/cid/"+content_id+".json"
   value := url.Values{}
   value["at"] = []string{at}
   req.URL.RawQuery = value.Encode()
   req.URL.Scheme = "https"
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
      panic(resp.Status)
   }
   file, err := os.Create(content_id + ".json")
   if err != nil {
      panic(err)
   }
   defer file.Close()
   file.ReadFrom(resp.Body)
}
