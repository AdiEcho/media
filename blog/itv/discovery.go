package main

import (
   "net/http"
   "net/url"
   "os"
)

// POST request not allowed
func main() {
   var req http.Request
   req.Header = http.Header{}
   req.URL = &url.URL{}
   req.URL.Host = "content-inventory.prd.oasvc.itv.com"
   req.URL.Path = "/discovery"
   req.URL.Scheme = "https"
   value := url.Values{}
   value["query"] = []string{query}
   req.URL.RawQuery = value.Encode()
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   resp.Write(os.Stdout)
}

const query = `query {
  brands(
    filter: {
      legacyId: "10/4008"
      features: [OUTBAND_WEBVTT, MPEG_DASH, WIDEVINE]
      available: "NOW"
      broadcaster: ITV
      platform: DOTCOM
      tiers: ["FREE", "PAID"]
    }
  ) {
    ccid
    title
    imageUrl(imageType: ITVX)
    numberOfAvailableSeries
    synopses {
      ninety
      epg
    }
    legacyId
    categories
    tier
    contentOwner
    genres {
      id
      name
    }
    visuallySigned
    partnership
    earliestAvailableSeries {
      earliestAvailableEpisode {
        broadcastDateTime
        seriesNumber
        episodeNumber
      }
    }
    latestAvailableSeries {
      longRunning
      fullSeries
      earliestAvailableEpisode {
        title
        broadcastDateTime
        seriesNumber
        episodeNumber
      }
      latestAvailableEpisode {
        title
        broadcastDateTime
        seriesNumber
        episodeNumber
      }
    }
    series(sortBy: SEQUENCE_ASC) {
      ccid
      seriesNumber
      legacyId
      fullSeries
      seriesType
      longRunning
      numberOfAvailableEpisodes
      latestAvailableEpisode {
        title
        broadcastDateTime
        seriesNumber
        episodeNumber
      }
      titles(sortBy: SEQUENCE_ASC) {
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
  }
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
  titlesToRedirect: titles(
    filter: {
      legacyId: "10/4008"
      titleTypes: [SPECIAL, FILM]
      available: "NOW"
      tiers: ["PAID", "FREE"]
    }
  ) {
    legacyId
    brandLegacyId
    title
    titleType
    brand {
      title
      legacyId
    }
  }
}`
