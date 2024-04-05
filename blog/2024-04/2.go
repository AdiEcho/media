package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.Header["Accept"] = []string{"application/vnd.companionservice.v1+json"}
   req.Header["Accept-Language"] = []string{"en-US,en;q=0.9"}
   req.Header["Content-Length"] = []string{"0"}
   req.Header["Content-Type"] = []string{"application/vnd.companionservice.v1+json"}
   req.Header["Origin"] = []string{"https://tv.clients.peacocktv.com"}
   req.Header["Referer"] = []string{"https://tv.clients.peacocktv.com/lightning/release/prod/android/5.4.13-ltv/chunk.worker.core.5cdfb0e4a520506acbf4.js"}
   req.Header["Sec-Fetch-Dest"] = []string{"empty"}
   req.Header["Sec-Fetch-Mode"] = []string{"cors"}
   req.Header["Sec-Fetch-Site"] = []string{"same-site"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Linux; Android 10; sdk_google_atv_x86 Build/QTU1.200805.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/83.0.4103.101 Mobile Safari/537.36"}
   req.Header["X-Requested-With"] = []string{"com.peacocktv.peacockandroid"}
   req.Header["X-Skyott-Activeterritory"] = []string{"US"}
   req.Header["X-Skyott-Bouquetid"] = []string{"6207143419709432117"}
   req.Header["X-Skyott-Device"] = []string{"TV"}
   req.Header["X-Skyott-Language"] = []string{"en"}
   req.Header["X-Skyott-Platform"] = []string{"ANDROIDTV"}
   req.Header["X-Skyott-Proposition"] = []string{"NBCUOTT"}
   req.Header["X-Skyott-Provider"] = []string{"NBCU"}
   req.Header["X-Skyott-Subbouquetid"] = []string{"0"}
   req.Header["X-Skyott-Territory"] = []string{"US"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "sas.peacocktv.com"
   req.URL.Scheme = "https"
   req.URL.Path = "/companion-service/journeys/957f6947-1fcc-4b3c-8bd8-ace430457eda"
   res, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res.Write(os.Stdout)
}
