package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "boot.pluto.tv"
   req.URL.Path = "/v4/start"
   req.URL.Scheme = "https"
   val := make(url.Values)
   
   req.Header["Accept"] = []string{"*/*"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.5"}
   req.Header["Origin"] = []string{"https://pluto.tv"}
   req.Header["Referer"] = []string{"https://pluto.tv/"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"same-site"}
   req.Header["Te"] = []string{"trailers"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"}
   req.Header["X-Forwarded-For"] = []string{"99.224.0.0"}
   val["appLaunchCount"] = []string{""}
   val["appName"] = []string{"web"}
   val["appVersion"] = []string{"8.0.0-111b2b9dc00bd0bea9030b30662159ed9e7c8bc6"}
   val["blockingMode"] = []string{""}
   val["clientID"] = []string{"766627e8-3c95-4cda-b743-f6d8bbd882aa"}
   val["clientModelNumber"] = []string{"1.0.0"}
   val["clientTime"] = []string{"2024-04-17T02:32:50.976Z"}
   val["deviceMake"] = []string{"firefox"}
   val["deviceModel"] = []string{"web"}
   val["deviceType"] = []string{"web"}
   val["deviceVersion"] = []string{"111.0.0"}
   val["drmCapabilities"] = []string{"widevine:L3"}
   val["episodeSlugs"] = []string{"ex-machina-2015-1-1-ptv1"}
   val["lastAppLaunchDate"] = []string{""}
   val["notificationVersion"] = []string{"1"}
   val["serverSideAds"] = []string{"false"}
   
   req.URL.RawQuery = val.Encode()
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
