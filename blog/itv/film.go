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
   resp.Write(os.Stdout)
}

// missing show
// missing year

const query = `
query {
  titles(
    filter: {
      legacyId: "10/3915/0002"
    }
  ) {
    ... on Episode {
      episodeNumber
      series {
        longRunning
        seriesNumber
        seriesType
      }
      contentOwner
      partnership
      versions {
        scheduleEvent {
          originalBroadcastDateTime
        }
      }
    }
    ... on Special {
      categories
      contentOwner
      episodeNumber
      partnership
      productionYear
      genres {
        id
        name
      }
      versions {
        scheduleEvent {
          originalBroadcastDateTime
        }
      }
    }
    ... on Film {
      categories
      contentOwner
      partnership
      productionYear
      genres {
        id
        name
        hubCategory
      }
    }
    ccid
    titleType
    legacyId
    brandLegacyId
    title
    channel {
      name
    }
    contentOwner
    partnership
    regionalisation
    broadcastDateTime
    imageUrl(imageType: ITVX)
    tier
    visuallySigned
    nextAvailableTitle {
      legacyId
    }
    series {
      fullSeries
      tier
      seriesNumber
      longRunning
    }
    brand {
      numberOfAvailableSeries
    }
    synopses {
      ninety
      epg
    }
    latestAvailableVersion {
      legacyId
      duration
      linearContent
      playlistUrl
      visuallySigned
      tier
      availability {
        start
        end
      }
      subtitled
      audioDescribed
      compliance {
        displayableGuidance
      }
      bsl {
        playlistUrl
      }
    }
  }
}
`
