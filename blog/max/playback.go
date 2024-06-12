package max

import (
   "io"
   "net/http"
   "net/url"
   "strings"
)

func playback() (*http.Response, error) {
   var req http.Request
   req.Header = make(http.Header)
   req.Method = "POST"
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "default.any-any.prd.api.max.com"
   req.URL.Path = "/any/playback/v1/playbackInfo"
   req.URL.Scheme = "https"
   req.Body = io.NopCloser(playback_body)
   req.Header["Content-Type"] = []string{"application/json"}
   req.Header["Cookie"] = []string{
      "st=eyJhbGciOiJSUzI1NiJ9.eyJqdGkiOiJ0b2tlbi1iODRmOTIxMy0yMzA0LTQ4MGEtOGEzMy1lOTViMTNhOGFiZjgiLCJpc3MiOiJmcGEtaXNzdWVyIiwic3ViIjoiVVNFUklEOmJvbHQ6Y2YwMWI0ZDItZDIyNS00Njc4LThkOTItOGU0NTg1MDhkN2U4IiwiaWF0IjoxNzE3OTcwODYxLCJleHAiOjIwMzMzMzA4NjEsInR5cGUiOiJBQ0NFU1NfVE9LRU4iLCJzdWJkaXZpc2lvbiI6ImJlYW1fYW1lciIsInNjb3BlIjoiZGVmYXVsdCIsInZlcnNpb24iOiJ2MiIsImFub255bW91cyI6ZmFsc2UsImRldmljZUlkIjoiMDQxMzdhYTItMWUxZS02ZjUyLTdhMDgtMTIyNDljODY0NjkwIn0.adU124rWw6-B55slVSnAn7gyd6wJA8sdWv-c2ayXkdrlGmXSRIosAnxf582ABO2ZCmguG0Lbm2S2ZlKMuRSwdT-QXfG8-EFW4LAaawiMc3xKuRn-uUmMCAhaewg_4TauEFdpAPDXAFOdO_wItNt7MoN1nQaW8C1Sa7jJzDpQhCDqv8DfEeZYfx_jQopnVyw6vUmz_W4m52wJAlmh_kW2fCJuUahywMKRHSBOBriBm1LL51gIOIcFxfLM3G6f-yb_ar9xqFqSIyaSpQ5Wj1t2T3IQPNh8gjrPw0O6A2zGo91S4reQHDB18kc-IzCFfGA2scGbLWZE7rcUpeZjK6eRiQ",
   }
   return http.DefaultClient.Do(&req)
}

var playback_body = strings.NewReader(`
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
