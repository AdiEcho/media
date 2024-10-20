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
   os.WriteFile("episode.json", dst.Bytes(), os.ModePerm)
}

const query = `
{
  versions(
    filter: {
      legacyId: "2_5460_0001.001"
      tiers: ["FREE", "PAID"]
      features: [OUTBAND_WEBVTT, MPEG_DASH, WIDEVINE]
      broadcaster: UNKNOWN
      platform: CTV
    }
  ) {
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
    variants {
      features
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
          categories
          genres {
            id
            name
          }
          title
        }
        nextAvailableTitle {
          latestAvailableVersion {
            legacyId
          }
        }
      }
      ... on Film {
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
