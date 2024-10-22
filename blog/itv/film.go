package main

import (
   "bytes"
   "encoding/json"
   "io"
   "net/http"
   "net/url"
   "os"
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
   var dst bytes.Buffer
   json.Indent(&dst, src, "", " ")
   os.WriteFile("film.json", dst.Bytes(), os.ModePerm)
}

const query = `
query {
  titles(
    filter: {
      titleTypes: [SPECIAL, FILM]
      brandLegacyId: "10/4008"
      available: "NOW"
      broadcaster: ITV
      platform: DOTCOM
      tiers: ["FREE", "PAID"]
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
      variants {
        features
        protection
      }
      bsl {
        playlistUrl
      }
    }
  }
}
`
