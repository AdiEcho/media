package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Header["Accept"] = []string{"application/json"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Content-Length"] = []string{"0"}
   req.Header["Origin"] = []string{"https://app.10ft.itv.com"}
   req.Header["Referer"] = []string{"https://app.10ft.itv.com/3.416.0/androidtv/programmes/10_4008/10a4008a0001"}
   req.Header["User-Agent"] = []string{"Mozilla/5.0 (Linux; Android 9; sdk_google_atv_x86 Build/PSR1.180720.121; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.158 Mobile Safari/537.36"}
   req.Header["X-Requested-With"] = []string{"air.ITVMobilePlayer"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = &url.URL{}
   req.URL.Host = "content-inventory.prd.oasvc.itv.com"
   req.URL.Path = "/discovery"
   value := url.Values{}
   value["query"] = []string{query}
   req.URL.RawQuery = value.Encode()
   req.URL.Scheme = "https"
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   src, err := io.ReadAll(resp.Body)
   if err != nil {
      panic(err)
   }
   dst := &bytes.Buffer{}
   json.Indent(dst, src, "", " ")
   fmt.Println(dst)
}

/*
itv.com/watch/community/10a3915/10a3915a0002
legacyId: "10/3915/0002"

itv.com/watch/pulp-fiction/10a3463
legacyId: "10/3463/0001"
*/
const query = `
query {
   titles(filter: {
      legacyId: "10/3915/0002"
   }) {
      brand {
         title
      }
      ... on Episode {
         seriesNumber
         episodeNumber
      }
      title
      ... on Film {
         productionYear
      }
   }
}
`
