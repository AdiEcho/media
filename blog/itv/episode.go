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
legacyId: "10/3915/0002#001"

itv.com/watch/pulp-fiction/10a3463
legacyId: "10/3463/0001#001"
*/
const query = `
{
   versions(filter: {
      legacyId: "10/3463/0001#001"
   }) {
      title {
         ... on Episode {
            brand {
               title
            }
            seriesNumber
            episodeNumber
         }
         title
         ... on Film {
            productionYear
         }
      }
   }
}
`
