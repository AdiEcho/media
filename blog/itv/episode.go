package main

import (
   "net/http"
   "net/url"
   "os"
)

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.Header["Accept"] = []string{"application/json"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Content-Length"] = []string{"0"}
   req.Header["Origin"] = []string{"https://app.10ft.itv.com"}
   req.Header["Referer"] = []string{"https://app.10ft.itv.com/3.416.0/androidtv/vod?productionId=2_5460_0001.001"}
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
   resp.Write(os.Stdout)
}

// itv.com/watch/community/10a3915/10a3915a0002
//legacyId: "10/3915/0002#001"

// itv.com/watch/pulp-fiction/10a3463
//legacyId: "10/3463/0001#001"

const query = `
{
   versions(filter: {
      legacyId: "10/3915/0002#001"
   }) {
      legacyId
      tier
      linearContent
      broadcastDateTime
      compliance {
         displayableGuidance
      }
      duration
      playlistUrl
      visuallySigned
      bsl {
         playlistUrl
      }
      availability {
         start
         end
      }
      title {
         ccid
         titleType
         title
         legacyId
         imageUrl(imageType: ITVX)
         synopses {
            ninety
            epg
         }
         series {
            fullSeries
            longRunning
            numberOfAvailableEpisodes
         }
         ... on Episode {
            brandLegacyId
            episodeNumber
            seriesNumber
            channel {
               name
            }
            brand {
               title
               categories
               genres {
                  id
                  name
               }
            }
            nextAvailableTitle {
               latestAvailableVersion {
                  legacyId
               }
            }
         }
         ... on Film {
            productionYear
            brandLegacyId
            title
            categories
            genres {
               id
               name
            }
            channel {
               name
            }
         }
         ... on Special {
            brandLegacyId
            title
            categories
            genres {
               id
               name
            }
            channel {
               name
            }
         }
      }
   }
}
`
