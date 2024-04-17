package main

import (
   "fmt"
   "io"
   "net/http"
   "net/url"
   "strings"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "boot.pluto.tv"
   req.URL.Path = "/v4/start"
   val := make(url.Values)
   val["appName"] = []string{"web"}
   val["appVersion"] = []string{"9"}
   val["clientID"] = []string{"9"}
   val["clientModelNumber"] = []string{"9"}
   val["drmCapabilities"] = []string{"widevine:L3"}
   val["episodeSlugs"] = []string{"ex-machina-2015-1-1-ptv1"}
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   text, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   fmt.Println(string(text))
   if strings.Contains(string(text), `/main.mpd"`) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
