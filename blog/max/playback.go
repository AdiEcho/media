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
   req.URL.Host = "default.any-any.prd.api.max.com"
   req.URL.Path = "/any/playback/v1/playbackInfo"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Cookie"] = []string{
      "st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi01ZjhmNDQzMy1jMWU2LTQ2ZjUtODA2NS04YjFiYTQxYzc1ZjAiLCJpc3MiOiJmcGEtaXNzdWVyIiwic3ViIjoiVVNFUklEOmJvbHQ6Y2YwMWI0ZDItZDIyNS00Njc4LThkOTItOGU0NTg1MDhkN2U4IiwiaWF0IjoxNzE3OTA0Mjg3LCJleHAiOjIwMzMyNjQyODcsInR5cGUiOiJBQ0NFU1NfVE9LRU4iLCJzdWJkaXZpc2lvbiI6ImJlYW1fYW1lciIsInNjb3BlIjoiZGVmYXVsdCIsInZlcnNpb24iOiJ2MiIsImFub255bW91cyI6ZmFsc2UsImRldmljZUlkIjoiZDI0MmEzOGQtMDY5My00ZTNjLTk2MDctYWU5ZTU0OWQwMDQwIn0.Nf0Oj1Q59IsoYQku3Wlow10rzQSjwCBtYdME7P6m7jyrQfjk3nb6lDJpb0vrdVbwidPngcXNQkxBS3681sobVdE1myBjrzUjLjIddztE1LgSHU818XxgXSS7YmB_GbwJbGM38MAk-uFi267yHqFZSR6FzSKnqcyfwKrss2D00k_4T-9Bw4BKymT_RNB_eDFO5SWvxLuOanglR7D-a64aTgujQlShxinMJBEPZ3qnM2KOLnZcdnXM1_b81qpSdY8goY9j-lk8a9JGSux8q0DhrgwXlGgMnaBn0JQQAe7kiW_1JbsIYmUipAbuo5VyT-0TCvEUtZYcVZoP5L9MIFni2g",
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
   "appBundle": "beam",
   "applicationSessionId": "b7804758-3377-4190-b429-ea7dee273880",
   "capabilities": {
      "manifests": {
         "formats": {
            "dash": {}
         }
      }
   },
   "consumptionType": "streaming",
   "deviceInfo": {
      "player": {
         "sdk": {
            "name": "Beam Player Desktop",
            "version": "4.1.0"
         },
         "mediaEngine": {
            "name": "GLUON_BROWSER",
            "version": "2.15.1"
         },
         "playerView": {
            "height": 864,
            "width": 1536
         }
      }
   },
   "editId": "1623fe4c-ef6e-4dd1-a10c-4a181f5f6579",
   "firstPlay": true,
   "gdpr": false,
   "playbackSessionId": "6ccb51f6-5cd6-4ce0-9ade-7f8a47e66474",
   "userPreferences": {}
}
`)
