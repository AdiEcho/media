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
   req.URL.Host = "cdnapisec.kaltura.com"
   req.URL.Path = "/api_v3/service/multirequest"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(body)
   req.Header["Content-Type"] = []string{"application/json"}
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}

var body = strings.NewReader(`
{
   "format": 1,
   "1": {
      "service": "session",
      "action": "startWidgetSession",
      "widgetId": "_2031841"
   },
   "2": {
      "service": "baseEntry",
      "action": "list",
      "ks": "{1:result:ks}",
      "filter": {
         "redirectFromEntryId": "1_kqvyiof1"
      }
   },
   "3": {
      "service": "baseEntry",
      "action": "getPlaybackContext",
      "entryId": "{2:result:objects:0:id}",
      "ks": "{1:result:ks}",
      "contextDataParams": {
         "flavorTags": "all"
      }
   }
}
`)
