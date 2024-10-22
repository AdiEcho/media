package main

import (
   "bytes"
   "encoding/json"
   "fmt"
   "io"
   "net/http"
   "net/url"
   "os"
)

//Show() string
//Season() int
//Episode() int
//Title() string
//Year() int

func main() {
   var req http.Request
   req.Header = http.Header{}
   req.URL = &url.URL{}
   req.URL.Host = "content-inventory.prd.oasvc.itv.com"
   req.URL.Path = "/discovery"
   req.URL.Scheme = "https"
   value := url.Values{}
   // episode
   value["query"] = []string{fmt.Sprintf(format, "2/5460/0023")}
   // film
   //value["query"] = []string{fmt.Sprintf(format, "10/4008/0001")}
   req.URL.RawQuery = value.Encode()
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   src, err := io.ReadAll(resp.Body)
   if err != nil {
      panic(err)
   }
   var dst bytes.Buffer
   json.Indent(&dst, src, "", " ")
   os.WriteFile("discovery.json", dst.Bytes(), os.ModePerm)
}

const format = `
query {
  titles(
    filter: { legacyId: %q }
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
